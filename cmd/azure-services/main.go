package azure_service

import (
	"context"
	"log"

	"github.com/microsoft/azure-devops-go-api/azuredevops"
	"github.com/microsoft/azure-devops-go-api/azuredevops/core"
)

func Authen() {
	organizationUrl := "https://dev.azure.com/myorg" // todo: replace value with your organization url
	personalAccessToken := "XXXXXXXXXXXXXXXXXXXXXXX" // todo: replace value with your PAT

	// Create a connection to your organization
	connection := azuredevops.NewPatConnection(organizationUrl, personalAccessToken)

	ctx := context.Background()

	// Create a client to interact with the Core area
	coreClient, err := core.NewClient(ctx, connection)
	if err != nil {
		log.Fatal(err)
	}

}
