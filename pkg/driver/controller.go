package driver

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	csi "github.com/container-storage-interface/spec/lib/go/csi"
	ctera "github.com/ctera/ctera-gateway-csi/pkg/ctera-openapi"
	"github.com/ctera/ctera-gateway-csi/pkg/driver/internal"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"k8s.io/klog"
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

type CteraVolumeId struct {
	filerAddress string
	shareName string
	path string
}

func (c *CteraVolumeId) toVolumeId() (*string, error) {
	bytes, err := json.Marshal(c)
	if err != nil {
		return nil, err
	}

	ret := string(bytes)
	return &ret, nil
}

func getCteraVolumeIdFromVolumeId(volumeId string) (*CteraVolumeId, error){
	var cteraVolumeId CteraVolumeId
	err := json.Unmarshal([]byte(volumeId), &cteraVolumeId)
	if err != nil {
		return nil, err
	}
	return &cteraVolumeId, nil
}

// controllerService represents the controller service of CSI driver
type controllerService struct {
	inFlight      *internal.InFlight
	driverOptions *DriverOptions
}

// newControllerService creates a new controller service
// it panics if failed to create the service
func newControllerService(driverOptions *DriverOptions) controllerService {
	return controllerService{
		inFlight:      internal.NewInFlight(),
		driverOptions: driverOptions,
	}
}

func (d *controllerService) CreateVolume(ctx context.Context, req *csi.CreateVolumeRequest) (*csi.CreateVolumeResponse, error) {
	klog.V(4).Infof("CreateVolume: called with args %+v", *req)
	shareName := req.GetName()
	if len(shareName) == 0 {
		return nil, status.Error(codes.InvalidArgument, "Name was not provided")
	}

	var (
		filerAddress string
		path string
	)

	for key, value := range req.GetParameters() {
		switch strings.ToLower(key) {
		case FilerAddressKey:
			filerAddress = value
		case PathKey:
			path = value
		default:
			return nil, status.Errorf(codes.InvalidArgument, "Invalid parameter key %s for CreateVolume", key)
		}
	}

	if len(path) == 0 {
		return nil, status.Error(codes.InvalidArgument, "Path was not provided")
	}

	if len(filerAddress) == 0 {
		return nil, status.Error(codes.InvalidArgument, "Filer address was not provided")
	}

	client, err := d.initClientConnection(ctx, filerAddress, req.GetSecrets())
	if err != nil {
		return nil, err
	}

	share, err := client.GetShareSafe(shareName)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	if share != nil {
		if d.canReuseShare(share, path) {
			createVolumeResponse, err := newCreateVolumeResponse(filerAddress, share)
			if err != nil {
				return nil, status.Error(codes.Internal, err.Error())
			}
			return createVolumeResponse, nil
		} else {
			return nil, status.Error(codes.AlreadyExists, "Share already exists with different parameters")
		}
	}

	share, err = client.CreateShare(shareName, path)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	createVolumeResponse, err := newCreateVolumeResponse(filerAddress, share)
	if err != nil {
		client.DeleteShareSafe(shareName)
		return nil, status.Error(codes.Internal, err.Error())
	}

	return createVolumeResponse, nil
}

func (d *controllerService) canReuseShare(share *ctera.Share, path string) (bool) {
	return share.GetDirectory() == path
}

func newCreateVolumeResponse(server string, share *ctera.Share) (*csi.CreateVolumeResponse, error) {
	cteraVolumeId := CteraVolumeId {
		filerAddress: server,
		shareName: share.GetName(),
		path: share.GetDirectory(),
	}

	volumeId, err := cteraVolumeId.toVolumeId()
	if err != nil {
		return nil, err
	}

	return &csi.CreateVolumeResponse{
		Volume: &csi.Volume{
			VolumeId: *volumeId,
			CapacityBytes: 0,
			VolumeContext: map[string]string{},
		},
	},
	nil
}

func (d *controllerService) DeleteVolume(ctx context.Context, req *csi.DeleteVolumeRequest) (*csi.DeleteVolumeResponse, error) {
	klog.V(4).Infof("DeleteVolume: called with args: %+v", *req)
	cteraVolumeId, err := getCteraVolumeIdFromVolumeId(req.GetVolumeId())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	client, err := d.initClientConnection(ctx, cteraVolumeId.filerAddress, req.GetSecrets())
	if err != nil {
		return nil, err
	}

	err = client.DeleteShareSafe(cteraVolumeId.shareName)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &csi.DeleteVolumeResponse{}, nil
}

func (d *controllerService) ControllerPublishVolume(ctx context.Context, req *csi.ControllerPublishVolumeRequest) (*csi.ControllerPublishVolumeResponse, error) {
	klog.V(4).Infof("ControllerPublishVolume: called with args %+v", *req)
	cteraVolumeId, err := getCteraVolumeIdFromVolumeId(req.GetVolumeId())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	client, err := d.initClientConnection(ctx, cteraVolumeId.filerAddress, req.GetSecrets())
	if err != nil {
		return nil, err
	}

	share, err := client.GetShareSafe(cteraVolumeId.shareName)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	if share == nil {
		return nil, status.Error(codes.NotFound, "Volume not found")
	}

	nodeAddress := "192.168.1.1"
	netmask := "255.255.0.0"
	perm := ctera.RW

	for _, trustedNfsClient := range share.GetTrustedNfsClients() {
		if trustedNfsClient.GetAddress() == nodeAddress && trustedNfsClient.GetNetmask() == netmask && trustedNfsClient.GetPerm() == perm {
			return &csi.ControllerPublishVolumeResponse{}, nil
		}
	}

	err = client.AddTrustedNfsClient(cteraVolumeId.shareName, nodeAddress, netmask, perm)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &csi.ControllerPublishVolumeResponse{}, nil
}

func (d *controllerService) ControllerUnpublishVolume(ctx context.Context, req *csi.ControllerUnpublishVolumeRequest) (*csi.ControllerUnpublishVolumeResponse, error) {
	klog.V(4).Infof("ControllerUnpublishVolume: called with args %+v", *req)
	cteraVolumeId, err := getCteraVolumeIdFromVolumeId(req.GetVolumeId())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	client, err := d.initClientConnection(ctx, cteraVolumeId.filerAddress, req.GetSecrets())
	if err != nil {
		return nil, err
	}

	share, err := client.GetShareSafe(cteraVolumeId.shareName)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	if share == nil {
		return nil, status.Error(codes.NotFound, "Volume not found")
	}

	nodeAddress := "192.168.1.1"
	netmask := "255.255.0.0"
	perm := ctera.RW

	for _, trustedNfsClient := range share.GetTrustedNfsClients() {
		if trustedNfsClient.GetAddress() == nodeAddress && trustedNfsClient.GetNetmask() == netmask && trustedNfsClient.GetPerm() == perm {
			err = client.RemoveTrustedNfsClient(cteraVolumeId.shareName, nodeAddress, netmask)
			if err != nil {
				return nil, status.Error(codes.Internal, err.Error())
			}
			break
		}
	}

	return &csi.ControllerUnpublishVolumeResponse{}, nil
}

func (d *controllerService) initClientConnection(ctx context.Context, filerAddress string, secrets map[string]string) (*CteraClient, error) {
	var (
		username string
		password string
	)

	for key, value := range secrets {
		switch strings.ToLower(key) {
		case FilerUsernameKey:
			username = value
		case FilerPasswordKey:
			password = value
		default:
			return nil, status.Errorf(codes.InvalidArgument, "Invalid secret key %s for CreateVolume", key)
		}
	}

	if len(username) == 0 {
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("Secret does not include %s", FilerUsernameKey))
	}

	if len(password) == 0 {
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("Secret does not include %s", FilerPasswordKey))
	}

	client, err := GetAuthenticatedCteraClient(filerAddress, username, password)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return client, nil
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

func (d *controllerService) GetCapacity(ctx context.Context, req *csi.GetCapacityRequest) (*csi.GetCapacityResponse, error) {
	klog.V(4).Infof("GetCapacity: called with args %+v", *req)
	return nil, status.Error(codes.Unimplemented, "")
}

func (d *controllerService) ListVolumes(ctx context.Context, req *csi.ListVolumesRequest) (*csi.ListVolumesResponse, error) {
	klog.V(4).Infof("ListVolumes: called with args %+v", *req)
	return nil, status.Error(codes.Unimplemented, "")
}

func (d *controllerService) ValidateVolumeCapabilities(ctx context.Context, req *csi.ValidateVolumeCapabilitiesRequest) (*csi.ValidateVolumeCapabilitiesResponse, error) {
	klog.V(4).Infof("ValidateVolumeCapabilities: called with args %+v", *req)
	return nil, status.Error(codes.Unimplemented, "Method not yet implemented")
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
