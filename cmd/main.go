package main

import (
	"flag"

	"github.com/ctera/ctera-gateway-csi/pkg/driver"

	"k8s.io/klog"
)

func main() {
	fs := flag.NewFlagSet("ctera-csi-driver", flag.ExitOnError)
	options := GetOptions(fs)

	drv, err := driver.NewDriver(
		driver.WithEndpoint(options.ServerOptions.Endpoint),
		driver.WithMode(options.DriverMode),
		driver.WithNodeIp(options.NodeOptions.NodeIp),
	)
	if err != nil {
		klog.Fatalln(err)
	}
	if err := drv.Run(); err != nil {
		klog.Fatalln(err)
	}
}
