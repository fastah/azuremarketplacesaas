# Go bindings for Azure Marketplace Saas Fulfilment and Metering APIs

This is an *Unofficial* Go SDK for Azure Marketplace Saas Fulfilment and Metering APIs, based on the OpenAPI 3 specification published by Microsoft's Marketplace team at this [repository](https://github.com/microsoft/commercial-marketplace-openapi/).

## Use at your own risk
This is not an official SDK and is not maintained by Microsoft. It is not guaranteed to be up to date with the latest changes in the Marketplace APIs. No guarantees are made about the correctness of the implementation.

## Installation

There are separate packages for Fulfilment and Metering APIs. You can install them using `go get`:

```go
import github.com/fastah/azuremarketplacesaas/fulfillment
```
```go
import github.com/fastah/azuremarketplacesaas/metering
```
If your SaaS product doesn't need metering, you can of course skip the metering package.


## Authorizing with Marketplace API endpoints

Due to limitations of the AutoRest code-generation tool that's been used to prepare this SDK, the code doesn't have any utilities for authorizing with the Marketplace API endpoints, as they are not the standard ARM endpoints. You will need to use the standard Azure SDKs to obtain an access token, and then use that token to call the Marketplace API endpoints.

## Fulfilment endpoint
See the [top-level guide to Fulfilment APIs](https://learn.microsoft.com/en-us/partner-center/marketplace/partner-center-portal/pc-saas-fulfillment-apis), and the two documents for [Subscription operations such as List, Activate](https://learn.microsoft.com/en-us/partner-center/marketplace/partner-center-portal/pc-saas-fulfillment-subscription-api) and [Operations such checking status of long-running tasks](https://learn.microsoft.com/en-us/partner-center/marketplace/partner-center-portal/pc-saas-fulfillment-operations-api).

## Metering endpoint
Carefully read Microsoft's [Marketplace metering service authentication strategies](https://learn.microsoft.com/en-us/partner-center/marketplace/marketplace-metering-service-authentication)

## Updating Go client from OpenAPI spec
### Submodule fetch from Microsoft's OpenAPI repo
```bash
git submodule update --init --recursive
```

### Install and Update AutoRest with Go extension
Follow the steps on the [AutoRest installation page](https://github.com/Azure/autorest/blob/main/docs/readme.md)

Install 
```bash
sudo npm install -g autorest
```
Purge all old Autorest and outdated extensions (careful!)
```bash
autorest --reset
```
Allow autorest to install latest Go extension
```bash
autorest --go --help
```

### Authentication to a Marketplace API endpoint

AutoRest's Go-specific generator doesn't appear to support generation of authorization code using `securityScheme` objects in the OpenAPI spec file, so we don't use any such capability: [Authentication configuration flags](https://github.com/Azure/autorest/blob/main/docs/generate/authentication.md). 