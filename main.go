package main

import (
	"context"
	"fmt"
	"os"
	// "strings"
	// "time"

	// "github.com/antihax/optional"
	ctera "github.com/ctera/ctera-gateway-csi/pkg/ctera"
)

func main() {
	// grab settings from environment
	address := os.Getenv("CTERA_GATEWAY_ADDRESS")
	username := os.Getenv("CTERA_GATEWAY_USER")
	password := os.Getenv("CTERA_GATEWAY_PASSWORD")

	// create golang client specific configuration
	configuration := ctera.NewConfiguration()
	configuration.Host = fmt.Sprintf("%s:9090", address)
	configuration.Servers = ctera.ServerConfigurations{
		{
			URL: fmt.Sprintf("http://%s:9090/v1.0", address),
			Description: "Main address",
		},
	}

	client := ctera.NewAPIClient(configuration)

	unauth := context.Background()
	credentials := *ctera.NewCredentials(username, password)

	jwt, _, err := client.LoginApi.LoginPost(unauth).Credentials(credentials).Execute()
	if err != nil {
		fmt.Println(err)

		os.Exit(1)
	}

	auth := context.WithValue(unauth, ctera.ContextAccessToken, jwt)

	searchFieldResults, _, err := client.UsersApi.UsersGet(auth).Execute()
	if err != nil {
		fmt.Println(err)

		os.Exit(1)
	}

	fmt.Println("Available search fields supported by ctera:")
	for _, field := range searchFieldResults {
		fmt.Println(fmt.Sprintf("%s", field.Username))
	}
}
