API_VERSION=$(shell cat api_version)
SDK_VERSION=$(shell cat sdk_version)
all: help

.PHONY: help
help:
	@echo "help:"
	@echo "- make gen  : regenerate SDK"
	@echo "- make test : run all tests"

.PHONY: gen
gen: clean osc

osc: osc-api/outscale.yaml
	rm -rf .sdk || true
	mkdir .sdk
	docker run -v $(PWD):/sdk --rm openapitools/openapi-generator-cli:v4.3.0 generate -i /sdk/osc-api/outscale.yaml -g go -c /sdk/gen.yml -o /sdk/.sdk --additional-properties=packageVersion=$(SDK_VERSION)
	mv .sdk osc

osc-api/outscale.yaml:
	git clone https://github.com/outscale/osc-api.git && cd osc-api && git checkout -b $(API_VERSION) $(API_VERSION)

.PHONY: clean
clean:
	rm -rf .sdk osc-api osc || true


.PHONY: test
test: build-examples reuse
	@echo all tests OK

.PHONY: reuse
reuse:
	docker run --volume $(PWD):/data fsfe/reuse:0.11.1 lint

.PHONY: build-examples
build-examples: examples
	find ./examples -type d -depth 1 -exec go build -o /dev/null {} \;

.PHONY: run-examples
run-examples:
	find ./examples -type d -depth 1 -exec go run {} \;

.PHONY: gofmt
gofmt:
	@find $(PWD)/examples/ -type f -name *.go -exec gofmt -l {} \;
	@find $(PWD)/examples/ -type f -name *.go -exec gofmt -w {} \;
