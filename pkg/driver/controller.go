/*
Copyright 2021, CTERA Networks.

Portions Copyright 2019 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package driver

import (
	"context"

	csi "github.com/container-storage-interface/spec/lib/go/csi"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	kubeclient "sigs.k8s.io/controller-runtime/pkg/client"

	ctera "github.com/ctera/ctera-gateway-openapi-go-client"
	"github.com/ctera/kubefiler-csi/pkg/driver/internal"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"k8s.io/klog"
)

const (
	pvcNameKey                   = "csi.storage.k8s.io/pvc/name"
	pvcNamespaceKey              = "csi.storage.k8s.io/pvc/namespace"
	kubefilerExportAnnotationKey = "kubefiler.ctera.com/kubefilerexport"
	secretNameSuffix             = "-kubefiler-credentials"
	gatewayUsernameKey           = "username"
	gatewayPasswordKey           = "password"
	serviceNameSuffix            = "-kubefiler"
)

var (
	// volumeCaps represents how the volume could be accessed.
	// It is SINGLE_NODE_WRITER since EBS volume could only be
	// attached to a single node at any given time.
	volumeCaps = []csi.VolumeCapability_AccessMode{
		{
			Mode: csi.VolumeCapability_AccessMode_MULTI_NODE_MULTI_WRITER,
		},
	}

	// controllerCaps represents the capability of controller service
	controllerCaps = []csi.ControllerServiceCapability_RPC_Type{
		csi.ControllerServiceCapability_RPC_CREATE_DELETE_VOLUME,
		csi.ControllerServiceCapability_RPC_PUBLISH_UNPUBLISH_VOLUME,
	}
)

// controllerService represents the controller service of CSI driver
type controllerService struct {
	kubeClient    kubeclient.Client
	inFlight      *internal.InFlight
	driverOptions *Options
}

// newControllerService creates a new controller service
// it panics if failed to create the service
func newControllerService(driverOptions *Options, kubeClient *kubeclient.Client) controllerService {
	return controllerService{
		kubeClient:    *kubeClient,
		inFlight:      internal.NewInFlight(),
		driverOptions: driverOptions,
	}
}

func (d *controllerService) CreateVolume(ctx context.Context, req *csi.CreateVolumeRequest) (*csi.CreateVolumeResponse, error) {
	klog.V(4).Infof("CreateVolume: called with args %+v", *req)
	pvcNamespace := req.GetParameters()[pvcNamespaceKey]
	pvcName := req.GetParameters()[pvcNameKey]

	pvc, err := d.getPvc(ctx, pvcNamespace, pvcName)
	if err != nil {
		return nil, err
	}

	klog.V(4).Infof("Got the Pvc: %+v", *pvc)

	associatedKubeFilerExport := pvc.GetAnnotations()[kubefilerExportAnnotationKey]

	klog.V(4).Infof("The associated KubeFilerExport's name is %s", associatedKubeFilerExport)

	_, err = internal.GetKubeFilerExport(ctx, d.kubeClient, pvcNamespace, associatedKubeFilerExport)
	if err != nil {
		return nil, err
	}

	createVolumeResponse, err := newCreateVolumeResponse(pvcNamespace, associatedKubeFilerExport)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return createVolumeResponse, nil
}

func (d *controllerService) getPvc(ctx context.Context, ns, name string) (*corev1.PersistentVolumeClaim, error) {
	pvc := &corev1.PersistentVolumeClaim{}
	err := d.kubeClient.Get(
		ctx,
		kubeclient.ObjectKey{
			Namespace: ns,
			Name:      name,
		},
		pvc,
	)

	return pvc, err
}

func newCreateVolumeResponse(namespace, kubeFilerExportName string) (*csi.CreateVolumeResponse, error) {
	cteraVolumeID := KubeFilerVolumeID{
		Namespace:           namespace,
		KubeFilerExportName: kubeFilerExportName,
	}

	volumeID, err := cteraVolumeID.ToVolumeID()
	if err != nil {
		return nil, err
	}

	return &csi.CreateVolumeResponse{
			Volume: &csi.Volume{
				VolumeId:      *volumeID,
				CapacityBytes: 0,
				VolumeContext: map[string]string{},
			},
		},
		nil
}

func (d *controllerService) DeleteVolume(ctx context.Context, req *csi.DeleteVolumeRequest) (*csi.DeleteVolumeResponse, error) {
	klog.V(4).Infof("DeleteVolume: called with args: %+v", *req)
	return &csi.DeleteVolumeResponse{}, nil
}

func (d *controllerService) ControllerPublishVolume(ctx context.Context, req *csi.ControllerPublishVolumeRequest) (*csi.ControllerPublishVolumeResponse, error) {
	klog.V(4).Infof("ControllerPublishVolume: called with args %+v", *req)
	kubeFilerVolumeID, err := getKubeFilerVolumeIDFromVolumeID(req.GetVolumeId())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	nodeAddress := req.GetNodeId()
	if len(nodeAddress) == 0 {
		return nil, status.Error(codes.InvalidArgument, "Node Id is empty")
	}

	filerAddress, filerUsername, filerPassword, err := d.getFilerLoginDetails(ctx, kubeFilerVolumeID.Namespace, kubeFilerVolumeID.KubeFilerExportName)
	if err != nil {
		return nil, err
	}

	client, err := GetAuthenticatedCteraClient(ctx, filerAddress, filerUsername, filerPassword)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	share, err := client.GetShareSafe(kubeFilerVolumeID.KubeFilerExportName)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	if share == nil {
		return nil, status.Error(codes.NotFound, "Volume not found")
	}

	netmask := "255.255.255.255"
	perm := ctera.RW

	for _, trustedNfsClient := range share.GetTrustedNfsClients() {
		if trustedNfsClient.GetAddress() == nodeAddress && trustedNfsClient.GetNetmask() == netmask && trustedNfsClient.GetPerm() == perm {
			return &csi.ControllerPublishVolumeResponse{}, nil
		}
	}

	err = client.AddTrustedNfsClient(kubeFilerVolumeID.KubeFilerExportName, nodeAddress, netmask, perm)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &csi.ControllerPublishVolumeResponse{}, nil
}

func (d *controllerService) ControllerUnpublishVolume(ctx context.Context, req *csi.ControllerUnpublishVolumeRequest) (*csi.ControllerUnpublishVolumeResponse, error) {
	klog.V(4).Infof("ControllerUnpublishVolume: called with args %+v", *req)
	kubeFilerVolumeID, err := getKubeFilerVolumeIDFromVolumeID(req.GetVolumeId())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	nodeAddress := req.GetNodeId()
	if len(nodeAddress) == 0 {
		return nil, status.Error(codes.InvalidArgument, "Node Id is empty")
	}

	filerAddress, filerUsername, filerPassword, err := d.getFilerLoginDetails(ctx, kubeFilerVolumeID.Namespace, kubeFilerVolumeID.KubeFilerExportName)
	if err != nil {
		if errors.IsNotFound(err) {
			klog.Warning("Filer details were not found - probably already deleted")
			return &csi.ControllerUnpublishVolumeResponse{}, nil
		}
		return nil, err
	}

	client, err := GetAuthenticatedCteraClient(ctx, filerAddress, filerUsername, filerPassword)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	share, err := client.GetShareSafe(kubeFilerVolumeID.KubeFilerExportName)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	if share != nil {
		netmask := "255.255.0.0"
		perm := ctera.RW

		for _, trustedNfsClient := range share.GetTrustedNfsClients() {
			if trustedNfsClient.GetAddress() == nodeAddress && trustedNfsClient.GetNetmask() == netmask && trustedNfsClient.GetPerm() == perm {
				err = client.RemoveTrustedNfsClient(kubeFilerVolumeID.KubeFilerExportName, nodeAddress, netmask)
				if err != nil {
					return nil, status.Error(codes.Internal, err.Error())
				}
				break
			}
		}
	}

	return &csi.ControllerUnpublishVolumeResponse{}, nil
}

func (d *controllerService) ControllerGetCapabilities(ctx context.Context, req *csi.ControllerGetCapabilitiesRequest) (*csi.ControllerGetCapabilitiesResponse, error) {
	klog.V(4).Infof("ControllerGetCapabilities: called with args %+v", *req)
	var caps []*csi.ControllerServiceCapability
	for _, cap := range controllerCaps {
		c := &csi.ControllerServiceCapability{
			Type: &csi.ControllerServiceCapability_Rpc{
				Rpc: &csi.ControllerServiceCapability_RPC{
					Type: cap,
				},
			},
		}
		caps = append(caps, c)
	}
	return &csi.ControllerGetCapabilitiesResponse{Capabilities: caps}, nil
}

func (d *controllerService) ValidateVolumeCapabilities(ctx context.Context, req *csi.ValidateVolumeCapabilitiesRequest) (*csi.ValidateVolumeCapabilitiesResponse, error) {
	klog.V(4).Infof("ValidateVolumeCapabilities: called with args %+v", *req)
	kubeFilerVolumeID, err := getKubeFilerVolumeIDFromVolumeID(req.GetVolumeId())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	volCaps := req.GetVolumeCapabilities()
	if len(volCaps) == 0 {
		return nil, status.Error(codes.InvalidArgument, "Volume capabilities not provided")
	}

	filerAddress, filerUsername, filerPassword, err := d.getFilerLoginDetails(ctx, kubeFilerVolumeID.Namespace, kubeFilerVolumeID.KubeFilerExportName)
	if err != nil {
		return nil, err
	}

	client, err := GetAuthenticatedCteraClient(ctx, filerAddress, filerUsername, filerPassword)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	share, err := client.GetShareSafe(kubeFilerVolumeID.KubeFilerExportName)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	if share == nil {
		return nil, status.Error(codes.NotFound, "Volume not found")
	}

	var confirmed *csi.ValidateVolumeCapabilitiesResponse_Confirmed
	if isValidVolumeCapabilities(volCaps) {
		confirmed = &csi.ValidateVolumeCapabilitiesResponse_Confirmed{VolumeCapabilities: volCaps}
	}
	return &csi.ValidateVolumeCapabilitiesResponse{
		Confirmed: confirmed,
	}, nil
}

func isValidVolumeCapabilities(volCaps []*csi.VolumeCapability) bool {
	hasSupport := func(cap *csi.VolumeCapability) bool {
		for _, c := range volumeCaps {
			if c.GetMode() == cap.AccessMode.GetMode() {
				return true
			}
		}
		return false
	}

	foundAll := true
	for _, c := range volCaps {
		if !hasSupport(c) {
			foundAll = false
		}
	}
	return foundAll
}

func (d *controllerService) getFilerLoginDetails(ctx context.Context, namespace, kubeFilerExportName string) (string, string, string, error) {
	kubeFilerExport, err := internal.GetKubeFilerExport(ctx, d.kubeClient, namespace, kubeFilerExportName)
	if err != nil {
		klog.Error(err, "Failed to get KubeFilerExport")
		return "", "", "", err
	}

	kubeFiler, err := internal.GetKubeFiler(ctx, d.kubeClient, namespace, kubeFilerExport.Spec.KubeFiler)
	if err != nil {
		klog.Error(err, "Failed to get KubeFiler")
		return "", "", "", err
	}

	kubeFilerSecret, err := internal.GetSecret(ctx, d.kubeClient, namespace, kubeFiler.GetName()+secretNameSuffix)
	if err != nil {
		klog.Error(err, "Failed to get Filer secret")
		return "", "", "", err
	}

	kubeFilerService, err := internal.GetService(ctx, d.kubeClient, namespace, kubeFiler.GetName()+serviceNameSuffix)
	if err != nil {
		klog.Error(err, "Failed to get Filer service")
		return "", "", "", err
	}

	return kubeFilerService.Spec.ClusterIP,
		string(kubeFilerSecret.Data[gatewayUsernameKey]),
		string(kubeFilerSecret.Data[gatewayPasswordKey]),
		nil
}

func (d *controllerService) GetCapacity(ctx context.Context, req *csi.GetCapacityRequest) (*csi.GetCapacityResponse, error) {
	klog.V(4).Infof("GetCapacity: called with args %+v", *req)
	return nil, status.Error(codes.Unimplemented, "")
}

func (d *controllerService) ListVolumes(ctx context.Context, req *csi.ListVolumesRequest) (*csi.ListVolumesResponse, error) {
	klog.V(4).Infof("ListVolumes: called with args %+v", *req)
	return nil, status.Error(codes.Unimplemented, "")
}

func (d *controllerService) ControllerExpandVolume(ctx context.Context, req *csi.ControllerExpandVolumeRequest) (*csi.ControllerExpandVolumeResponse, error) {
	klog.V(4).Infof("ControllerExpandVolume: called with args %+v", *req)
	return nil, status.Error(codes.Unimplemented, "")
}

func (d *controllerService) ControllerGetVolume(ctx context.Context, req *csi.ControllerGetVolumeRequest) (*csi.ControllerGetVolumeResponse, error) {
	klog.V(4).Infof("ControllerGetVolume: called with args %+v", *req)
	return nil, status.Error(codes.Unimplemented, "")
}

func (d *controllerService) CreateSnapshot(ctx context.Context, req *csi.CreateSnapshotRequest) (*csi.CreateSnapshotResponse, error) {
	klog.V(4).Infof("CreateSnapshot: called with args %+v", req)
	return nil, status.Error(codes.Unimplemented, "")
}

func (d *controllerService) DeleteSnapshot(ctx context.Context, req *csi.DeleteSnapshotRequest) (*csi.DeleteSnapshotResponse, error) {
	klog.V(4).Infof("DeleteSnapshot: called with args %+v", req)
	return nil, status.Error(codes.Unimplemented, "")
}

func (d *controllerService) ListSnapshots(ctx context.Context, req *csi.ListSnapshotsRequest) (*csi.ListSnapshotsResponse, error) {
	klog.V(4).Infof("ListSnapshots: called with args %+v", req)
	return nil, status.Error(codes.Unimplemented, "")
}
