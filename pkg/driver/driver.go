/*
Copyright 2021, CTERA Networks.

Portions Copyright 2019 The Kubernetes Authors.

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

package driver

import (
	"context"
	"fmt"
	"net"

	csi "github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/ctera/kubefiler-csi/pkg/util"
	kubefilerv1alpha1 "github.com/ctera/kubefiler-operator/api/v1alpha1"
	"google.golang.org/grpc"
	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	"k8s.io/klog"
	ctrl "sigs.k8s.io/controller-runtime"
	kubeclient "sigs.k8s.io/controller-runtime/pkg/client"
)

// Mode is the operating mode of the CSI driver.
type Mode string

const (
	// ControllerMode is the mode that only starts the controller service.
	ControllerMode Mode = "controller"
	// NodeMode is the mode that only starts the node service.
	NodeMode Mode = "node"
	// AllMode is the mode that only starts both the controller and the node service.
	AllMode Mode = "all"
)

const (
	driverName = "csi.kubefiler.ctera.com"
)

type Driver struct {
	controllerService
	nodeService

	srv     *grpc.Server
	options *Options
}

type Options struct {
	endpoint                     string
	mode                         Mode
	nodeIP                       string
	KubeFilerOperatorNameSpace   string
	KubeFilerLockerConfigMapName string
}

func NewDriver(options ...func(*Options)) (*Driver, error) {
	klog.V(4).Infof("Driver: %v Version: %v", driverName, driverVersion)

	driverOptions := Options{
		endpoint: DefaultCSIEndpoint,
		mode:     AllMode,
	}
	for _, option := range options {
		option(&driverOptions)
	}

	if err := validateDriverOptions(&driverOptions); err != nil {
		return nil, fmt.Errorf("invalid driver options: %v", err)
	}

	kubeClient, err := initKubeClient()
	if err != nil {
		return nil, err
	}

	driver := Driver{
		options: &driverOptions,
	}

	switch driverOptions.mode {
	case ControllerMode:
		driver.controllerService = newControllerService(&driverOptions, kubeClient)
	case NodeMode:
		driver.nodeService = newNodeService(&driverOptions, kubeClient)
	case AllMode:
		driver.controllerService = newControllerService(&driverOptions, kubeClient)
		driver.nodeService = newNodeService(&driverOptions, kubeClient)
	default:
		return nil, fmt.Errorf("unknown mode: %s", driverOptions.mode)
	}

	return &driver, nil
}

func (d *Driver) Run() error {
	scheme, addr, err := util.ParseEndpoint(d.options.endpoint)
	if err != nil {
		return err
	}

	listener, err := net.Listen(scheme, addr)
	if err != nil {
		return err
	}

	logErr := func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		resp, err := handler(ctx, req)
		if err != nil {
			klog.Errorf("GRPC error: %v", err)
		}
		return resp, err
	}
	opts := []grpc.ServerOption{
		grpc.UnaryInterceptor(logErr),
	}
	d.srv = grpc.NewServer(opts...)

	csi.RegisterIdentityServer(d.srv, d)

	switch d.options.mode {
	case ControllerMode:
		csi.RegisterControllerServer(d.srv, d)
	case NodeMode:
		csi.RegisterNodeServer(d.srv, d)
	case AllMode:
		csi.RegisterControllerServer(d.srv, d)
		csi.RegisterNodeServer(d.srv, d)
	default:
		return fmt.Errorf("unknown mode: %s", d.options.mode)
	}

	klog.V(4).Infof("Listening for connections on address: %#v", listener.Addr())
	return d.srv.Serve(listener)
}

func (d *Driver) Stop() {
	d.srv.Stop()
}

func WithEndpoint(endpoint string) func(*Options) {
	return func(o *Options) {
		o.endpoint = endpoint
	}
}

func WithMode(mode Mode) func(*Options) {
	return func(o *Options) {
		o.mode = mode
	}
}

func WithNodeIP(nodeIP string) func(*Options) {
	return func(o *Options) {
		o.nodeIP = nodeIP
	}
}

func WithKubeFilerOperatorNameSpace(kubeFilerOperatorNameSpace string) func(*Options) {
	return func(o *Options) {
		o.KubeFilerOperatorNameSpace = kubeFilerOperatorNameSpace
	}
}

func WithKubeFilerLockerConfigMapName(kubeFilerLockerConfigMapName string) func(*Options) {
	return func(o *Options) {
		o.KubeFilerLockerConfigMapName = kubeFilerLockerConfigMapName
	}
}

func validateDriverOptions(options *Options) error {
	if err := validateMode(options.mode); err != nil {
		return fmt.Errorf("invalid mode: %v", err)
	}

	return nil
}

func validateMode(mode Mode) error {
	if mode != AllMode && mode != ControllerMode && mode != NodeMode {
		return fmt.Errorf("mode is not supported (actual: %s, supported: %v)", mode, []Mode{AllMode, ControllerMode, NodeMode})
	}

	return nil
}

func initKubeClient() (*kubeclient.Client, error) {
	scheme := runtime.NewScheme()
	utilruntime.Must(clientgoscheme.AddToScheme(scheme))
	utilruntime.Must(kubefilerv1alpha1.AddToScheme(scheme))

	kubeClient, err := kubeclient.New(
		ctrl.GetConfigOrDie(),
		kubeclient.Options{
			Scheme: scheme,
		},
	)
	if err != nil {
		klog.Error("Failed to initialize the kubeclient")
		return nil, err
	}

	return &kubeClient, err
}
