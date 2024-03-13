# AutoRest for Marketplace SaaS Fulfillment API code generation

This markdown file can be passed as `autorest fulfillment.md`, and the block below will be automagically parsed as configuration for the AutoRest code generation step. 

``` yaml
go: true
input-file: '../commercial-marketplace-openapi/Microsoft.Marketplace.SaaS/2018-08-31/saasapi.v2.json'
output-folder: '../fulfillment'
file-prefix: 'f_'
module-version: 0.1.0
module: azurempsaas-fulfillment
openapi-type: data-plane
azure-arm: true
license-header: MICROSOFT_MIT_NO_VERSION
payload-flattening-threshold: 3
credential-scope: "20e940b3-4c77-4b0b-9a53-9e16a1b010a7/.default"
```
