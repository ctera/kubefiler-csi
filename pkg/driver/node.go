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
	"fmt"
	"os"
	"strings"

	csi "github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/ctera/kubefiler-csi/pkg/driver/internal"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"k8s.io/klog"
	mount "k8s.io/mount-utils"
	kubeclient "sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	sharePathBase = "/nfs/shares"
)

var (
	// nodeCaps represents the capability of node service.
	nodeCaps = []csi.NodeServiceCapability_RPC_Type{}
)

// nodeService represents the node service of CSI driver
type nodeService struct {
	kubeClient    kubeclient.Client
	mounter       mount.Interface
	inFlight      *internal.InFlight
	driverOptions *Options
}

// newNodeService creates a new node service
func newNodeService(driverOptions *Options, kubeClient *kubeclient.Client) nodeService {
	return nodeService{
		kubeClient:    *kubeClient,
		mounter:       mount.New(""),
		inFlight:      internal.NewInFlight(),
		driverOptions: driverOptions,
	}
}

func (ns *nodeService) NodePublishVolume(ctx context.Context, req *csi.NodePublishVolumeRequest) (*csi.NodePublishVolumeResponse, error) {
	klog.V(4).Infof("NodePublishVolume: called with args %+v", *req)
	targetPath := req.GetTargetPath()
	if len(targetPath) == 0 {
		klog.Error("Target path not provided")
		return nil, status.Error(codes.InvalidArgument, "Target path not provided")
	}

	notMnt, err := ns.mounter.IsLikelyNotMountPoint(targetPath)
	if err != nil {
		if os.IsNotExist(err) {
			if err := os.MkdirAll(targetPath, 0750); err != nil {
				klog.Error(err, "Failed to create the directory")
				return nil, status.Error(codes.Internal, err.Error())
			}
			notMnt = true
		} else {
			klog.Error(err, "Failed to evaluate path")
			return nil, status.Error(codes.Internal, err.Error())
		}
	}
	if !notMnt {
		klog.Error("Path exists but is already a mount point")
		return &csi.NodePublishVolumeResponse{}, nil
	}

	kubeFilerVolumeID, err := getKubeFilerVolumeIDFromVolumeID(req.GetVolumeId())
	if err != nil {
		klog.Error(err, "Failed to get KubeFilerVolumeID")
		return nil, status.Error(codes.Internal, err.Error())
	}

	if req.GetVolumeCapability() == nil {
		klog.Error(err, "Failed to get VolumeCapability")
		return nil, status.Error(codes.InvalidArgument, "Volume capability missing in request")
	}

	filerAddress, err := ns.getFilerAddress(ctx, kubeFilerVolumeID)
	if err != nil {
		klog.Error(err, "Failed to get Filer Address")
		return nil, err
	}
	sharePath := sharePathBase + "/" + kubeFilerVolumeID.KubeFilerExportName
	source := fmt.Sprintf("%s:%s", filerAddress, sharePath)

	mountOptions := req.GetVolumeCapability().GetMount().GetMountFlags()
	mountOptions = append(mountOptions, "nolock")
	if req.GetReadonly() {
		mountOptions = append(mountOptions, "ro")
	}

	klog.V(2).Infof("NodePublishVolume: volumeID(%v) source(%s) targetPath(%s) mountflags(%v)", kubeFilerVolumeID, source, targetPath, mountOptions)
	err = ns.mounter.Mount(source, targetPath, "nfs", mountOptions)
	if err != nil {
		if os.IsPermission(err) {
			klog.Error("Failed to mount due to permissions")
			return nil, status.Error(codes.PermissionDenied, err.Error())
		}
		if strings.Contains(err.Error(), "invalid argument") {
			klog.Error("Failed to mount due to an invalid argument")
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
		klog.Error(err, "Failed to mount due to other reason")
		return nil, status.Error(codes.Internal, err.Error())
	}

	klog.V(4).Info("NodePublishVolume: Done successfully")
	return &csi.NodePublishVolumeResponse{}, nil
}

func (ns *nodeService) NodeUnpublishVolume(ctx context.Context, req *csi.NodeUnpublishVolumeRequest) (*csi.NodeUnpublishVolumeResponse, error) {
	klog.V(4).Infof("NodeUnpublishVolume: called with args %+v", *req)
	kubeFilerVolumeID, err := getKubeFilerVolumeIDFromVolumeID(req.GetVolumeId())
	if err != nil {
		klog.Error(err, "Failed to get KubeFilerVolumeID")
		return nil, status.Error(codes.Internal, err.Error())
	}

	targetPath := req.GetTargetPath()
	if len(targetPath) == 0 {
		klog.Error("Target path not provided")
		return nil, status.Error(codes.InvalidArgument, "Target path missing in request")
	}

	notMnt, err := ns.mounter.IsLikelyNotMountPoint(targetPath)
	if err != nil {
		if os.IsNotExist(err) {
			klog.Error("Target path not found")
			return nil, status.Error(codes.NotFound, "Targetpath not found")
		}
		klog.Error(err, "Failed to evalute path")
		return nil, status.Error(codes.Internal, err.Error())
	}
	if notMnt {
		klog.Error("Target path is not mounted")
		return &csi.NodeUnpublishVolumeResponse{}, nil
	}

	klog.V(2).Infof("NodeUnpublishVolume: CleanupMountPoint %s on volumeID(%+v)", targetPath, kubeFilerVolumeID)
	err = mount.CleanupMountPoint(targetPath, ns.mounter, false)
	if err != nil {
		klog.Error(err, "Failed to unmount")
		return nil, status.Error(codes.Internal, err.Error())
	}

	klog.V(4).Info("NodeUnpublishVolume done successfully")
	return &csi.NodeUnpublishVolumeResponse{}, nil
}

func (ns *nodeService) NodeGetCapabilities(ctx context.Context, req *csi.NodeGetCapabilitiesRequest) (*csi.NodeGetCapabilitiesResponse, error) {
	klog.V(4).Infof("NodeGetCapabilities: called with args %+v", *req)
	var caps []*csi.NodeServiceCapability
	for _, cap := range nodeCaps {
		c := &csi.NodeServiceCapability{
			Type: &csi.NodeServiceCapability_Rpc{
				Rpc: &csi.NodeServiceCapability_RPC{
					Type: cap,
				},
			},
		}
		caps = append(caps, c)
	}
	return &csi.NodeGetCapabilitiesResponse{Capabilities: caps}, nil
}

func (ns *nodeService) NodeGetInfo(ctx context.Context, req *csi.NodeGetInfoRequest) (*csi.NodeGetInfoResponse, error) {
	klog.V(4).Infof("NodeGetInfo: called with args %+v", *req)
	return &csi.NodeGetInfoResponse{
		NodeId: ns.driverOptions.nodeIP,
	}, nil
}

func (ns *nodeService) getFilerAddress(ctx context.Context, kubeFilerVolumeID *KubeFilerVolumeID) (string, error) {
	kubeFilerExport, err := internal.GetKubeFilerExport(ctx, ns.kubeClient, kubeFilerVolumeID.Namespace, kubeFilerVolumeID.KubeFilerExportName)
	if err != nil {
		klog.Error(err, "Failed to get KubeFilerExport")
		return "", err
	}

	kubeFiler, err := internal.GetKubeFiler(ctx, ns.kubeClient, kubeFilerVolumeID.Namespace, kubeFilerExport.Spec.KubeFiler)
	if err != nil {
		klog.Error(err, "Failed to get KubeFiler")
		return "", err
	}

	kubeFilerService, err := internal.GetService(ctx, ns.kubeClient, kubeFilerVolumeID.Namespace, kubeFiler.GetName()+serviceNameSuffix)
	if err != nil {
		klog.Error(err, "Failed to get Filer service")
		return "", err
	}

	return kubeFilerService.Spec.ClusterIP, nil
}

func (ns *nodeService) NodeGetVolumeStats(ctx context.Context, req *csi.NodeGetVolumeStatsRequest) (*csi.NodeGetVolumeStatsResponse, error) {
	klog.V(4).Infof("NodeGetVolumeStats: called with args %+v", *req)
	return nil, status.Error(codes.Unimplemented, "Method not yet implemented")
}

func (ns *nodeService) NodeStageVolume(ctx context.Context, req *csi.NodeStageVolumeRequest) (*csi.NodeStageVolumeResponse, error) {
	klog.V(4).Infof("NodeStageVolume: called with args %+v", *req)
	return nil, status.Error(codes.Unimplemented, "Method not yet implemented")
}

func (ns *nodeService) NodeUnstageVolume(ctx context.Context, req *csi.NodeUnstageVolumeRequest) (*csi.NodeUnstageVolumeResponse, error) {
	klog.V(4).Infof("NodeUnstageVolume: called with args %+v", *req)
	return nil, status.Error(codes.Unimplemented, "Method not yet implemented")
}

func (ns *nodeService) NodeExpandVolume(ctx context.Context, req *csi.NodeExpandVolumeRequest) (*csi.NodeExpandVolumeResponse, error) {
	klog.V(4).Infof("NodeExpandVolume: called with args %+v", *req)
	return nil, status.Error(codes.Unimplemented, "Method not yet implemented")
}
