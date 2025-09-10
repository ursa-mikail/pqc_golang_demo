.PHONY: test test-race bench cover clean build run help

# Default target
help:
	@echo "Available targets:"
	@echo "  test      - Run all tests"
	@echo "  test-race - Run tests with race detection"
	@echo "  bench     - Run benchmarks"
	@echo "  cover     - Generate coverage report"
	@echo "  build     - Build demo application"
	@echo "  run       - Run demo application"
	@echo "  clean     - Clean generated files"

# Check that all imports work
check-imports:
	@echo "Checking imports..."
	go run tools/check_imports.go

# Run all tests
test:
	@echo "Running all tests..."
	go test -v ./...

# Run both import check and tests
test-all: check-imports test
	@echo "All checks completed!"

# Run tests with race detection
test-race:
	@echo "Running tests with race detection..."
	go test -race -v ./...

# Run benchmarks
bench:
	@echo "Running benchmarks..."
	go test -bench=. -benchmem ./...

# Generate coverage report
cover:
	@echo "Generating coverage report..."
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

# Test specific packages
test-util:
	go test -v ./util

test-ciphering:
	go test -v ./ciphering

test-signing:
	go test -v ./signing

test-hashing:
	go test -v ./hashing

# Benchmark specific packages
bench-ciphering:
	go test -bench=. -benchmem ./cipher

bench-signing:
	go test -bench=. -benchmem ./signing

bench-hashing:
	go test -bench=. -benchmem ./hashing

# Build demo application
build:
	@echo "Building demo application..."
	go build -o pqc_demo main.go

# Run demo application
run: build
	@echo "Running demo application..."
	./pqc_demo

# Performance profiling
profile-cpu:
	go test -bench=BenchmarkEncapsulate -cpuprofile=cpu.prof ./cipher
	go tool pprof cpu.prof

profile-mem:
	go test -bench=BenchmarkSign -memprofile=mem.prof ./signing
	go tool pprof mem.prof

# Clean generated files
clean:
	@echo "Cleaning generated files..."
	rm -f coverage.out coverage.html pqc_demo *.prof

# Initialize project (for new installations)
init:
	@echo "Initializing project..."
	go mod tidy
	@echo "Project ready!"

# Run security checks
security:
	@echo "Running security checks..."
	go list -json -m all | nancy sleuth
	gosec ./...

# Format code
fmt:
	@echo "Formatting code..."
	go fmt ./...

# Run linter
lint:
	@echo "Running linter..."
	golangci-lint run

# Full CI pipeline
ci: fmt lint test-race cover
	@echo "CI pipeline completed successfully!"