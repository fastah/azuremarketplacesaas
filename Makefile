# These files are input OpenAPI specs from the Microsoft-supplied submodule.
SPECFILE_METERING=./commercial-marketplace-openapi/Microsoft.Marketplace.Metering/2018-08-31/meteringapi.v1
SPECFILE_FULFILLMENT=./commercial-marketplace-openapi/Microsoft.Marketplace.SaaS/2018-08-31/saasapi.v2

# Use this to nuke your Autorest installation and install latest tool and language plugins
autorest-go-update-with-reset: 
	# You need to prefix sudo for the install globally command below.
	sudo npm install -g autorest
	autorest --reset
	autorest --go --help

test-metering:
	source env.sh && go test -v metering/*.go -run=TestMeteringClient

# Builds both the metering and fulfillment clients
build: codegen-metering codegen-fulfillment

codegen-metering: metering.md
	autorest metering.md --input-file=$(SPECFILE_METERING).json
	cd ./metering && go mod tidy && go get -u ./... && go mod tidy

codegen-fulfillment: fulfillment.md
	autorest fulfillment.md --input-file=$(SPECFILE_FULFILLMENT).json
	cd ./fulfillment && go mod tidy && go get -u ./... && go mod tidy
