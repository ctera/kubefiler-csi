package options

import (
	"flag"
	"testing"
)

func TestServerOptions(t *testing.T) {
	testCases := []struct {
		name  string
		flag  string
		found bool
	}{
		{
			name:  "lookup desired flag",
			flag:  "endpoint",
			found: true,
		},
		{
			name:  "fail for non-desired flag",
			flag:  "some-other-flag",
			found: false,
		},
	}

	for _, tc := range testCases {
		flagSet := flag.NewFlagSet("test-flagset", flag.ContinueOnError)
		serverOptions := &ServerOptions{}

		t.Run(tc.name, func(t *testing.T) {
			serverOptions.AddFlags(flagSet)

			flag := flagSet.Lookup(tc.flag)
			found := flag != nil
			if found != tc.found {
				t.Fatalf("result not equal\ngot:\n%v\nexpected:\n%v", found, tc.found)
			}
		})
	}
}
