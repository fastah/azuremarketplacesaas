package azurempsaas

import (
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/fastah/azurempsaas/fulfillment"
)

// Hello returns a greeting.
func NewFulfillmentClient() string {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("Authentication failure: %+v", err)
	}

	// Azure Resource Management clients accept the credential as a parameter
	client, _ := arm.NewClient("azurempsaas", string(fulfillment.APIVersionTwoThousandEighteen0831), cred, &arm.ClientOptions{})

	log.Print("Authenticated to subscription", client)
}
