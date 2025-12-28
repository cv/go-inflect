.PHONY: test bench bench-save bench-compare lint

# Run tests
test:
	go test -v ./...

# Run benchmarks
bench:
	go test -bench=. -benchmem -count=6 ./...

# Save benchmark baseline
bench-save:
	go test -bench=. -benchmem -count=6 ./... > benchmarks/baseline.txt

# Compare against baseline
bench-compare:
	go test -bench=. -benchmem -count=6 ./... > benchmarks/new.txt
	go run golang.org/x/perf/cmd/benchstat benchmarks/baseline.txt benchmarks/new.txt

# Run linter
lint:
	golangci-lint run
