module github.com/ctera/ctera-gateway-csi

go 1.15

require (
	github.com/container-storage-interface/spec v1.4.0
	github.com/ctera/ctera-gateway-csi/pkg/ctera-openapi v1.0.0
	github.com/golang/protobuf v1.4.3 // indirect
	github.com/google/go-cmp v0.5.2 // indirect
	golang.org/x/net v0.0.0-20210224082022-3d97a244fca7 // indirect
	golang.org/x/sys v0.0.0-20210225134936-a50acf3fe073 // indirect
	golang.org/x/text v0.3.4 // indirect
	google.golang.org/grpc v1.31.0
	k8s.io/klog v1.0.0
	k8s.io/mount-utils v0.21.1
	k8s.io/utils v0.0.0-20210527160623-6fdb442a123b // indirect
)

replace github.com/ctera/ctera-gateway-csi/pkg/ctera-openapi v1.0.0 => ./pkg/ctera-openapi
