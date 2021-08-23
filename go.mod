module github.com/ctera/kubefiler-csi

go 1.15

require (
	github.com/container-storage-interface/spec v1.4.0
	github.com/ctera/ctera-gateway-openapi-go-client v0.0.0-20210823130300-9b654e0e43ca
	github.com/ctera/kubefiler-operator v0.0.0-20210823143308-4fa51b24ca42
	golang.org/x/net v0.0.0-20210224082022-3d97a244fca7 // indirect
	golang.org/x/sys v0.0.0-20210225134936-a50acf3fe073 // indirect
	google.golang.org/grpc v1.31.0
	k8s.io/api v0.20.2
	k8s.io/apimachinery v0.20.2
	k8s.io/client-go v0.20.2
	k8s.io/klog v1.0.0
	k8s.io/mount-utils v0.21.1
	k8s.io/utils v0.0.0-20210527160623-6fdb442a123b // indirect
	sigs.k8s.io/controller-runtime v0.8.3
)
