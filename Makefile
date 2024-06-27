SHELL := bash
.SHELLFLAGS := -eu -o pipefail -c
.DEFAULT_GOAL := help

.PHONY: lint fmt test test-side lint-side fmt-side

## Lint the code
lint:
	@npx mega-linter-runner --flavor go

lin-side: ## Lint the code exp. MK_DESC_POSITION=side
	@npx mega-linter-runner --flavor go

## Format the code
fmt:
	@go fmt ./...

fmt-side: ## Format the code exp. MK_DESC_POSITION=side
	@go fmt ./...


## Run the tests
test: 
	@go clean -testcache
	@go test -v ./... && echo -e "\033[32mOK\033[0m" || echo -e "\033[31mERROR\033[0m";

test-side: ## Run the tests exp. MK_DESC_POSITION=side
	@go clean -testcache
	@go test -v ./... && echo -e "\033[32mOK\033[0m" || echo -e "\033[31mERROR\033[0m";
