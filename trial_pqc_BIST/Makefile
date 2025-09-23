# PQC BIST Demo Makefile

.PHONY: all build test clean run deps help

# Default target
all: build

# Build the application
build:
	@echo "Building PQC BIST Demo..."
	go build -o pqc_bist_demo main.go

# Run the application
run:
	@echo "Running PQC BIST Demo..."
	go mod tidy
	go run main.go

# Install dependencies
deps:
	@echo "Installing dependencies..."
	go mod tidy
	go mod download

# Initialize go module if not exists
init:
	@if [ ! -f go.mod ]; then \
		echo "Initializing Go module..."; \
		go mod init pqc_bist_demo; \
	fi

# Run tests
test:
	@echo "Running tests..."
	go test -v ./...

# Run benchmarks
bench:
	@echo "Running benchmarks..."
	go test -bench=. -benchmem

# Clean build artifacts
clean:
	@echo "Cleaning..."
	rm -f pqc_bist_demo
	rm -rf pqc_*.json
	go clean

# Generate test vectors
vectors:
	@echo "Generating test vectors..."
	go run main.go > test_vectors_output.txt
	@echo "Test vectors saved to test_vectors_output.txt"

# Run with verbose output
verbose:
	@echo "Running with verbose output..."
	go run -v main.go

# Check for security vulnerabilities
security:
	@echo "Checking for security vulnerabilities..."
	go list -json -m all | nancy sleuth

# Format code
fmt:
	@echo "Formatting code..."
	go fmt ./...

# Lint code
lint:
	@echo "Linting code..."
	golangci-lint run

# Show help
help:
	@echo "Available targets:"
	@echo "  build    - Build the application"
	@echo "  run      - Run the application"
	@echo "  deps     - Install dependencies"
	@echo "  init     - Initialize Go module"
	@echo "  test     - Run tests"
	@echo "  bench    - Run benchmarks"
	@echo "  clean    - Clean build artifacts"
	@echo "  vectors  - Generate test vectors to file"
	@echo "  verbose  - Run with verbose output"
	@echo "  fmt      - Format code"
	@echo "  lint     - Lint code"
	@echo "  help     - Show this help"