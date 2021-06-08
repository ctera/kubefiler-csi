package options

import (
	"flag"
)

// ControllerOptions contains options and configuration settings for the controller service.
type ControllerOptions struct {
	// ID of the kubernetes cluster.
	KubernetesClusterID string
}

func (s *ControllerOptions) AddFlags(fs *flag.FlagSet) {
	fs.StringVar(&s.KubernetesClusterID, "k8s-tag-cluster-id", "", "ID of the Kubernetes cluster used for tagging provisioned EBS volumes (optional).")
}
