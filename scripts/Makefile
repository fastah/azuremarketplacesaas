
CODEGEN_DIR=../
GO_PKGNAME=azurempsaas
USE_GOAUTOREST=@autorest/go@4.0.0-preview.51

build: codegen-metering codegen-fulfillment
	
codegen-metering:
	autorest metering.md --module-version="0.0.1" --module="metering" --output-folder=$(CODEGEN_DIR)/metering --use=$(USE_GOAUTOREST)
	rm $(CODEGEN_DIR)/metering/go.mod


codegen-fulfillment:
	autorest fulfillment.md --module-version="0.0.1" --module="fulfillment" --output-folder=$(CODEGEN_DIR)/fulfillment --use=$(USE_GOAUTOREST)
	rm $(CODEGEN_DIR)/fulfillment/go.mod
