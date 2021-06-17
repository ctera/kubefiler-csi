/*
Copyright 2021, CTERA Networks.

Portions Copyright 2020 The Kubernetes Authors.

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
