module github.com/ctera/ctera-gateway-csi

go 1.15

require (
	github.com/antihax/optional v1.0.0
	github.com/ctera/ctera-gateway-csi/pkg/ctera-openapi v1.0.0
)

replace github.com/ctera/ctera-gateway-csi/pkg/ctera-openapi v1.0.0 => ./pkg/ctera-openapi
