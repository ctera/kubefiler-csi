package options

import (
	"flag"

	"github.com/ctera/ctera-gateway-csi/pkg/driver"
)

// ServerOptions contains options and configuration settings for the driver server.
type ServerOptions struct {
	// Endpoint is the endpoint that the driver server should listen on.
	Endpoint string
	// HttpEndpoint is the endpoint that the HTTP server for metrics should listen on.
	HttpEndpoint string
}

func (s *ServerOptions) AddFlags(fs *flag.FlagSet) {
	fs.StringVar(&s.Endpoint, "endpoint", driver.DefaultCSIEndpoint, "Endpoint for the CSI driver server")
	fs.StringVar(&s.HttpEndpoint, "http-endpoint", "", "The TCP network address where the HTTP server for metrics will listen (example: `:8080`). The default is empty string, which means the server is disabled.")
}
