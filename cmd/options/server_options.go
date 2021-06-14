package options

import (
	"flag"

	"github.com/ctera/ctera-gateway-csi/pkg/driver"
)

// ServerOptions contains options and configuration settings for the driver server.
type ServerOptions struct {
	// Endpoint is the endpoint that the driver server should listen on.
	Endpoint string
}

func (s *ServerOptions) AddFlags(fs *flag.FlagSet) {
	fs.StringVar(&s.Endpoint, "endpoint", driver.DefaultCSIEndpoint, "Endpoint for the CSI driver server")
}
