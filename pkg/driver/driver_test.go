package driver

import (
	"testing"
)

func TestWithEndpoint(t *testing.T) {
	value := "endpoint"
	options := &DriverOptions{}
	WithEndpoint(value)(options)
	if options.endpoint != value {
		t.Fatalf("expected endpoint option got set to %q but is set to %q", value, options.endpoint)
	}
}

func TestWithMode(t *testing.T) {
	value := Mode("mode")
	options := &DriverOptions{}
	WithMode(value)(options)
	if options.mode != value {
		t.Fatalf("expected mode option got set to %q but is set to %q", value, options.mode)
	}
}

func TestWithClusterID(t *testing.T) {
	var id string = "test-cluster-id"
	options := &DriverOptions{}
	WithKubernetesClusterID(id)(options)
	if options.kubernetesClusterID != id {
		t.Fatalf("expected kubernetesClusterID option got set to %s but is set to %s", id, options.kubernetesClusterID)
	}
}
