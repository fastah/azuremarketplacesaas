build: codegen-metering codegen-fulfillment

autorest-go-update-with-reset: 
	# You need to prefix sudo for the install globally command below.
	npm install -g autorest
	autorest --reset
	autorest --go --help

codegen-metering:
	autorest metering.md
	cd ../metering && go mod tidy && go get -u ./... && go mod tidy

codegen-fulfillment:
	autorest fulfillment.md
	cd ../fulfillment && go mod tidy && go get -u ./... && go mod tidy
