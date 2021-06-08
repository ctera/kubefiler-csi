package options

import (
	"flag"
	"testing"
)

func TestNodeOptions(t *testing.T) {
	testCases := []struct {
		name  string
		flag  string
		found bool
	}{
		{
			name:  "lookup desired flag",
			flag:  "volume-attach-limit",
			found: true,
		},
		{
			name:  "fail for non-desired flag",
			flag:  "some-flag",
			found: false,
		},
	}

	for _, tc := range testCases {
		flagSet := flag.NewFlagSet("test-flagset", flag.ContinueOnError)
		nodeOptions := &NodeOptions{}

		t.Run(tc.name, func(t *testing.T) {
			nodeOptions.AddFlags(flagSet)

			flag := flagSet.Lookup(tc.flag)
			found := flag != nil
			if found != tc.found {
				t.Fatalf("result not equal\ngot:\n%v\nexpected:\n%v", found, tc.found)
			}
		})
	}
}
