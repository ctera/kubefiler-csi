module github.com/ctera/ctera-gateway-csi

go 1.15

require (
	github.com/container-storage-interface/spec v1.4.0
	github.com/ctera/ctera-gateway-csi/pkg/ctera-openapi v1.0.0
	github.com/golang/protobuf v1.4.3
	github.com/kubernetes-csi/csi-proxy/client v0.2.2
	github.com/stretchr/testify v1.6.1
	golang.org/x/sys v0.0.0-20210225134936-a50acf3fe073
	google.golang.org/grpc v1.31.0
	k8s.io/component-base v0.21.1
	k8s.io/klog v1.0.0
	k8s.io/klog/v2 v2.8.0
	k8s.io/mount-utils v0.21.1
	k8s.io/utils v0.0.0-20210527160623-6fdb442a123b
)

replace github.com/ctera/ctera-gateway-csi/pkg/ctera-openapi v1.0.0 => ./pkg/ctera-openapi
