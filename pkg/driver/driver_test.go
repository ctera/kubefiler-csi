package driver

import (
	"testing"
)

func TestWithEndpoint(t *testing.T) {
	value := "endpoint"
	options := &Options{}
	WithEndpoint(value)(options)
	if options.endpoint != value {
		t.Fatalf("expected endpoint option got set to %q but is set to %q", value, options.endpoint)
	}
}

func TestWithMode(t *testing.T) {
	value := Mode("mode")
	options := &Options{}
	WithMode(value)(options)
	if options.mode != value {
		t.Fatalf("expected mode option got set to %q but is set to %q", value, options.mode)
	}
}
