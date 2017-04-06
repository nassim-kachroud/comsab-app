.PHONY: install
install:
	go get github.com/mattn/goveralls

.PHONY: build
build:
	CGO_ENABLED=0 go build -a -ldflags '-s' -installsuffix cgo -o after-sales-invoicing ./src

.PHONY: test-unit
test-unit:
	chmod 777 -R script/code-coverage
	script/code-coverage --coveralls

.PHONY: test-integration
test-integration:
	go test -v ./src/github.com/ricardo-ch/after-sales-invoicing/src/api/... -tags=integration