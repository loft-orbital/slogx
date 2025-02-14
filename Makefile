# Copyright © Loft Orbital Solutions Inc.
# Use of this source code is governed by a Apache-2.0-style
# license that can be found in the LICENSE file.

all: mocks build test lint ## Run all targets
.PHONY: all

build: ## Check if the library compiles
	@echo "Checking if library compiles..."
	go build -o /dev/null ./...
.PHONY: build

test: ## Run tests
	@echo "Running tests..."
	go test -coverprofile=cover.out  -v ./...
.PHONY: test

lint: ## Lint the code
	@echo "Linting the code..."
	golangci-lint run ./...
.PHONY: lint

fmt: ## Format the code
	@echo "Formatting code..."
	gofmt -s -w .
.PHONY: fmt

tidy: ## Tidy up Go modules
	@echo "Tidying up modules..."
	go mod tidy
.PHONY: tidy

mocks: ## Generate mocks using mockery
	@echo "Generating mocks..."
	mockery
.PHONY: mocks

help: ## Show help
	@grep -E '^[0-9a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
.PHONY: help
