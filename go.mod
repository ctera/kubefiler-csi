module github.com/ctera/kubefiler-csi

go 1.15

require (
	github.com/container-storage-interface/spec v1.4.0
	github.com/ctera/ctera-gateway-openapi-go-client v0.3.1
	github.com/ctera/kubefiler-operator v0.2.0
	github.com/pborman/uuid v1.2.1
	github.com/werf/lockgate v0.0.0-20211004100849-f85d5325b201
	google.golang.org/grpc v1.38.0
	k8s.io/api v0.22.2
	k8s.io/apimachinery v0.22.2
	k8s.io/client-go v0.22.2
	k8s.io/klog v1.0.0
	k8s.io/mount-utils v0.21.1
	sigs.k8s.io/controller-runtime v0.10.1
)
