SHELL := bash
.SHELLFLAGS := -eu -o pipefail -c
.DEFAULT_GOAL := help

.PHONY: test fmt

## Run tests
test:
	go test ./...

## Format source code
fmt:
	go fmt ./...
