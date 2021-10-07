GOBIN := $(shell go env GOPATH)/bin

models:
	rm -f apimodel/*_generated.go
	swagger generate model --spec=apimodel/swagger_apimodel.yml --target=. --model-package=apimodel --config-file=apimodel/swagger_apimodel_config.yml

.PHONY: lint
lint: ${GOBIN}/golangci-lint
	golangci-lint run -v

.PHONY: test
test: lint
	go test -v -count=1 -failfast ./...