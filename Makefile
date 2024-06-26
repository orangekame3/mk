SHELL := bash
.SHELLFLAGS := -eu -o pipefail -c
.DEFAULT_GOAL := help

.PHONY: lint fmt help

lint: ## Lint the code
	@npx mega-linter-runner --flavor go

fmt: ## Format the code
	@go fmt ./...

test: ## Run the tests
	@go clean -testcache
	@go test -v ./... && echo -e "\033[32mOK\033[0m" || echo -e "\033[31mERROR\033[0m";

help: ## Show this help message
	@echo "Usage: make [target]"
	@echo ""
	@echo "Available targets:"
	@echo ""
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(filter-out .env,$(MAKEFILE_LIST)) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
