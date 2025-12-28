.PHONY: help deps build test lint bench bench-save bench-compare

.DEFAULT_GOAL := help

help: ## Print help message
	@grep -E '^[\/a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

deps: ## Download dependencies
	go mod download

build: ## Build the project
	go build ./...

test: ## Run tests with race detection and coverage
	go test -v -race -coverprofile=coverage.out ./...

lint: ## Run linter
	golangci-lint run

bench: ## Run benchmarks
	go test -bench=. -benchmem -count=6 ./...

bench-save: ## Save benchmark baseline
	go test -bench=. -benchmem -count=6 ./... > benchmarks/baseline.txt

bench-compare: ## Compare against baseline
	go test -bench=. -benchmem -count=6 ./... > benchmarks/new.txt
	go run golang.org/x/perf/cmd/benchstat benchmarks/baseline.txt benchmarks/new.txt
