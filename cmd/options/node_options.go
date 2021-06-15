package options

import (
	"flag"
)

// NodeOptions contains options and configuration settings for the node service.
type NodeOptions struct {
	// NodeIP is the IP of the Node in which the node service is running on
	// It is used for the Trusted NFS Clients list
	// Users should pass it via "status.hostIP"
	NodeIP string
}

func (o *NodeOptions) AddFlags(fs *flag.FlagSet) {
	fs.StringVar(&o.NodeIP, "node-ip", "", "IP Address if the host on which the service is running")
}
