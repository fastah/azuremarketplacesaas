package metering

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

// The Marketplace metering service application ID (20e940b3-4c77-4b0b-9a53-9e16a1b010a7).
// See https://learn.microsoft.com/en-us/partner-center/marketplace-offers/marketplace-metering-service-authentication
const AzureMarketplaceWellKnownTenantID = "20e940b3-4c77-4b0b-9a53-9e16a1b010a7"

// NewMeteringClient creates a new instance of OperationsClient with the specified values.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewMeteringClient(credential azcore.TokenCredential, scopes []string, options *policy.ClientOptions) (*OperationsClient, error) {
	popts := runtime.PipelineOptions{
		PerCall: []policy.Policy{runtime.NewBearerTokenPolicy(credential, scopes, nil)}}
	cl, err := azcore.NewClient("UnofficialMeteringGo", "v0.1.0", popts, options)
	if err != nil {
		return nil, err
	}
	soc := &OperationsClient{internal: cl}
	return soc, nil
}
