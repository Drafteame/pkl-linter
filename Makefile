# Define default goal
.DEFAULT_GOAL := help

.PHONY: help test lint format build

help: ## Show this help message
	@echo "Usage: make [target] tf_dir=<path> [workspace=dev]\n"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

test: ## Execute tests
	go test ./... -race

lint: ## Lint the code
	golangci-lint run

format: ## Format the code
	goimports-reviser -format -excludes=".git,.github" ./...

build: ## Build the code
	go build -ldflags="-s -w -X github.com/Drafteame/pkl-linter/pkg/version.Version=v0.0.0" -o ./.bin/pkl-linter ./main.go
