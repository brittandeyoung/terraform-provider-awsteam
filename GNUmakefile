PKG_NAME            ?= internal
TESTARGS                ?= "-run=TestAcc"

default: testacc

# Please keep targets in alphabetical order

build: ## Build provider
	go install

gen:
	go generate

golangci-lint: ## Lint Go source (via golangci-lint)
	@echo "==> Checking source code with golangci-lint..."
	@golangci-lint run \
		--config .golangci.yml \
		./$(PKG_NAME)/...

test:
	go test ./...

# Run acceptance tests
testacc:
	TF_ACC=1 go test ./... -v $(TESTARGS) -timeout 120m

.PHONY: 
	- build
	- generate
	- golangci-lint
	- testacc
	- test