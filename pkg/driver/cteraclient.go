package driver

import (
	"context"
	"fmt"

	"k8s.io/klog"
	ctera "github.com/ctera/ctera-gateway-csi/pkg/ctera-openapi"
)

type CteraClientError struct {
	error string
}

// Error returns non-empty string if there was an error.
func (e CteraClientError) Error() string {
	return e.error
}

type CteraClient struct {
	client *ctera.APIClient
	ctx *context.Context
}

func GetAuthenticatedCteraClient(filerAddress, username, password string) (*CteraClient, error) {
	client, err := NewCteraClient(filerAddress)
	if err != nil {
		return nil, err
	}

	err = client.Authenticate(username, password)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func NewCteraClient(filerAddress string) (*CteraClient, error) {
	configuration := ctera.NewConfiguration()
	configuration.Host = fmt.Sprintf("%s:9090", filerAddress)
	configuration.Servers = ctera.ServerConfigurations{
		{
			URL: fmt.Sprintf("http://%s:9090/v1.0", filerAddress),
			Description: "Main address",
		},
	}

	return &CteraClient{
		client: ctera.NewAPIClient(configuration),
		ctx: nil,
	}, nil
}

func (c *CteraClient) Authenticate(username, password string) (error) {
	if c.ctx != nil {
		return CteraClientError{
			error: "Already authenticated",
		}
	}

	unauth := context.Background()
	credentials := ctera.NewCredentials(username, password)

	jwt, _, err := c.client.LoginApi.LoginPost(unauth).Credentials(*credentials).Execute()
	if err != nil {
		klog.Error(err)
		return err
	}

	auth := context.WithValue(unauth, ctera.ContextAccessToken, jwt)
	c.ctx = &auth

	return nil
}

func (c *CteraClient) GetShareSafe(name string)(*ctera.Share, error) {
	share, _, err := c.client.SharesApi.SharesNameGet(*c.ctx, name).Execute()
	if err != nil {
		if c.getStatusFromError(err) != 404 {
			return nil, err
		}
		return nil, nil
	}
	return &share, nil
}

func (c *CteraClient) CreateShare(name, path string)(*ctera.Share, error) {
	share := ctera.NewShare(name)
	share.Directory = &path
	// TODO Do we need to override any default parameters

	_, err := c.client.SharesApi.SharesPost(*c.ctx).Share(*share).Execute()
	if err != nil {
		return nil, err
	}

	return c.GetShareSafe(name)
}

func (c *CteraClient) DeleteShareSafe(name string)(error) {
	_, err := c.client.SharesApi.SharesNameDelete(*c.ctx, name).Execute()
	if err != nil {
		if c.getStatusFromError(err) != 404 {
			return err
		}
	}
	return nil
}

func (c *CteraClient) AddTrustedNfsClient(shareName, address, netmask string, perm ctera.FileAccessMode) (error) {
	trustedNfsClients := []ctera.NFSv3AccessControlEntry{*ctera.NewNFSv3AccessControlEntry(address, netmask, perm)}
	_, err := c.client.SharesApi.CteraGatewayOpenapiApiSharesAddTrustedNfsClients(*c.ctx, shareName).NFSv3AccessControlEntry(trustedNfsClients).Execute()
	return err
}

func (c *CteraClient) RemoveTrustedNfsClient(shareName, address, netmask string) (error) {
	_, err := c.client.SharesApi.CteraGatewayOpenapiApiSharesRemoveTrustedNfsClient(*c.ctx, shareName).Address(address).Netmask(netmask).Execute()
	return err
}

func (c *CteraClient) getStatusFromError(err error) (int32) {
	genericOpenAPIError, ok := err.(ctera.GenericOpenAPIError)
	if !ok {
		return -1
	}

	errorMessage, ok := genericOpenAPIError.Model().(ctera.ErrorMessage)
	if !ok {
		return -1
	}

	return errorMessage.GetStatus()
}
