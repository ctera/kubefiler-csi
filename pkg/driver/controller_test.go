package driver

import (
	"context"
	"reflect"
	"testing"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"github.com/ctera/ctera-gateway-csi/pkg/driver/internal"
)

func TestNewControllerService(t *testing.T) {
	testCases := []struct {
		name    string
		options *Options
	}{
		{
			name: "Working",
			options: &Options{
				endpoint: "test",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			controllerSvc := newControllerService(tc.options)

			if !reflect.DeepEqual(controllerSvc.driverOptions, tc.options) {
				t.Fatalf("expected driverOptions attribute to be equal to input")
			}
		})
	}
}

func TestCreateVolume(t *testing.T) {
	testCases := []struct {
		name         string
		req          *csi.CreateVolumeRequest
		res          *csi.CreateVolumeResponse
		expectedCode codes.Code
	}{
		{
			name: "Fail with No Name",
			req: &csi.CreateVolumeRequest{
				Name: "",
			},
			res: nil,
			expectedCode: codes.InvalidArgument,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			service := controllerService{
				inFlight:      internal.NewInFlight(),
				driverOptions: &Options{},
			}
			ctx := context.Background()

			res, err := service.CreateVolume(ctx, tc.req)
			if tc.expectedCode != codes.OK && (err != nil || status.Code(err) != tc.expectedCode) {
				if err == nil {
					t.Errorf("The operation succeeded while it was supposed to fail with err %d", tc.expectedCode)
				}
				actualCode := status.Code(err)
				if actualCode != tc.expectedCode {
					t.Errorf("The operation failed with the code %d instead of %d", actualCode, tc.expectedCode)
				}
				return
			}

			if !reflect.DeepEqual(res, tc.res) {
				t.Fatalf("Response is not as expected")
			}

		})
	}
}