package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/ctera/ctera-gateway-csi/cmd/options"
	"github.com/ctera/ctera-gateway-csi/pkg/driver"

	"k8s.io/klog"
)

// Options is the combined set of options for all operating modes.
type Options struct {
	DriverMode driver.Mode

	*options.ServerOptions
	*options.ControllerOptions
	*options.NodeOptions
}

// used for testing
var osExit = os.Exit

// GetOptions parses the command line options and returns a struct that contains
// the parsed options.
func GetOptions(fs *flag.FlagSet) *Options {
	var (
		version = fs.Bool("version", false, "Print the version and exit.")

		args = os.Args[1:]
		mode = driver.AllMode

		serverOptions     = options.ServerOptions{}
		controllerOptions = options.ControllerOptions{}
		nodeOptions       = options.NodeOptions{}
	)

	serverOptions.AddFlags(fs)
	klog.InitFlags(fs)

	if len(os.Args) > 1 {
		cmd := os.Args[1]

		switch {
		case cmd == string(driver.ControllerMode):
			controllerOptions.AddFlags(fs)
			args = os.Args[2:]
			mode = driver.ControllerMode

		case cmd == string(driver.NodeMode):
			nodeOptions.AddFlags(fs)
			args = os.Args[2:]
			mode = driver.NodeMode

		case cmd == string(driver.AllMode):
			controllerOptions.AddFlags(fs)
			nodeOptions.AddFlags(fs)
			args = os.Args[2:]

		case strings.HasPrefix(cmd, "-"):
			controllerOptions.AddFlags(fs)
			nodeOptions.AddFlags(fs)
			args = os.Args[1:]

		default:
			fmt.Printf("unknown command: %s: expected %q, %q or %q", cmd, driver.ControllerMode, driver.NodeMode, driver.AllMode)
			os.Exit(1)
		}
	}

	if err := fs.Parse(args); err != nil {
		panic(err)
	}

	if *version {
		info, err := driver.GetVersionJSON()
		if err != nil {
			klog.Fatalln(err)
		}
		fmt.Println(info)
		osExit(0)
	}

	return &Options{
		DriverMode: mode,

		ServerOptions:     &serverOptions,
		ControllerOptions: &controllerOptions,
		NodeOptions:       &nodeOptions,
	}
}
