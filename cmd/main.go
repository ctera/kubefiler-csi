package main

import (
	"flag"
	"net/http"

	"github.com/ctera/ctera-gateway-csi/pkg/driver"
	"k8s.io/component-base/metrics/legacyregistry"

	"k8s.io/klog"
)

func main() {
	fs := flag.NewFlagSet("ctera-csi-driver", flag.ExitOnError)
	options := GetOptions(fs)

	drv, err := driver.NewDriver(
		driver.WithEndpoint(options.ServerOptions.Endpoint),
		driver.WithExtraTags(options.ControllerOptions.ExtraTags),
		driver.WithExtraVolumeTags(options.ControllerOptions.ExtraVolumeTags),
		driver.WithMode(options.DriverMode),
		driver.WithVolumeAttachLimit(options.NodeOptions.VolumeAttachLimit),
		driver.WithKubernetesClusterID(options.ControllerOptions.KubernetesClusterID),
		driver.WithAwsSdkDebugLog(options.ControllerOptions.AwsSdkDebugLog),
	)
	if err != nil {
		klog.Fatalln(err)
	}
	if err := drv.Run(); err != nil {
		klog.Fatalln(err)
	}
}
