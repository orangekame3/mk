SHELL := bash
.SHELLFLAGS := -eu -o pipefail -c
.DEFAULT_GOAL := help

.PHONY: test fmt

test: ## Run tests
	export MK_DESC_POSITION=side && go test ./...

fmt: ## Format source code
	export MK_DESC_POSITION=side && go fmt ./...
