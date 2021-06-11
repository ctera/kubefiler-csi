package options

import (
	"flag"
)

// NodeOptions contains options and configuration settings for the node service.
type NodeOptions struct {
	// VolumeAttachLimit specifies the value that shall be reported as "maximum number of attachable volumes"
	// in CSINode objects. It is similar to https://kubernetes.io/docs/concepts/storage/storage-limits/#custom-limits
	// which allowed administrators to specify custom volume limits by configuring the kube-scheduler. Also, each AWS
	// machine type has different volume limits. By default, the EBS CSI driver parses the machine type name and then
	// decides the volume limit. However, this is only a rough approximation and not good enough in most cases.
	// Specifying the volume attach limit via command line is the alternative until a more sophisticated solution presents
	// itself (dynamically discovering the maximum number of attachable volume per EC2 machine type, see also
	// https://github.com/kubernetes-sigs/aws-ebs-csi-driver/issues/347).
	VolumeAttachLimit int64
	// NodeIp is the IP of the Node in which the node service is running on
	// It is used for the Trusted NFS Clients list
	// Users should pass it via "status.hostIP"
	NodeIp string
}

func (o *NodeOptions) AddFlags(fs *flag.FlagSet) {
	fs.Int64Var(&o.VolumeAttachLimit, "volume-attach-limit", -1, "Value for the maximum number of volumes attachable per node. If specified, the limit applies to all nodes. If not specified, the value is approximated from the instance type.")
	fs.StringVar(&o.NodeIp, "node-ip", "", "IP Address if the host on which the service is running")
}
