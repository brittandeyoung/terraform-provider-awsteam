TESTARGS                ?= "-run=TestAcc"

default: testacc

gen:
	go generate

test:
	go test ./...

# Run acceptance tests
testacc:
	TF_ACC=1 go test ./... -v $(TESTARGS) -timeout 120m

.PHONY: 
	- generate
	- testacc
	- test