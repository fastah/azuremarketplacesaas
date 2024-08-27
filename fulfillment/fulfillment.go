package fulfillment

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

//const azureMPWellKnownTenantID = "20e940b3-4c77-4b0b-9a53-9e16a1b010a7"

// NewFulfillmentClient creates a new instance of SubscriptionOperationsClient with the specified values.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewFulfillmentClient(credential azcore.TokenCredential, options *policy.ClientOptions) (*SubscriptionOperationsClient, error) {
	popts := runtime.PipelineOptions{
		PerCall: []policy.Policy{runtime.NewBearerTokenPolicy(credential, []string{}, nil)}}
	cl, err := azcore.NewClient("UnofficialFulfillmentGo", "0.1.0", popts, options)
	if err != nil {
		return nil, err
	}
	soc := &SubscriptionOperationsClient{internal: cl}
	return soc, nil
}
