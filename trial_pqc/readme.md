
## Basic Testing
```
# Run all tests
go test ./...

# Run specific package tests
go test ./ciphering
go test ./signing
go test ./hashing
go test ./util

# Run with verbose output
go test -v ./ciphering
```

## Coverage Testing: Generate coverage report
```
go test -cover ./...
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out -o coverage.html
```

## Benchmark Testing
```
# Run all benchmarks
go test -bench=. ./...

# Run specific benchmarks with memory stats
go test -bench=BenchmarkEncapsulate -benchmem ./cipher
go test -bench=BenchmarkSign -benchmem ./signing
```

# Using Makefile
```
make test      # Run all tests
make bench     # Run benchmarks
make cover     # Generate coverage report
make run       # Build and run demo
make clean     # Clean generated files
```

# Go Test Commands Guide

## Basic Test Commands

### Run All Tests
```bash
# Run all tests in the project
go test ./...

# Run tests with verbose output
go test -v ./...

# Run tests with coverage
go test -cover ./...

# Run tests with detailed coverage report
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out -o coverage.html
```

### Run Tests by Package
```bash
# Test specific packages
go test ./util
go test ./ciphering  
go test ./signing
go test ./hashing

# With verbose output
go test -v ./ciphering
```

### Run Specific Tests
```bash
# Run tests matching a pattern
go test -run TestGenerateKeyPair ./cipher
go test -run TestHash ./hashing
go test -run TestSign ./signing

# Run tests with verbose output
go test -v -run TestEncapsulateDecapsulate ./ciphering
```

## Benchmark Commands

### Run All Benchmarks
```bash
# Run all benchmarks
go test -bench=. ./...

# Run benchmarks with memory stats
go test -bench=. -benchmem ./...

# Run benchmarks multiple times for accuracy
go test -bench=. -count=5 ./...
```

### Run Specific Benchmarks
```bash
# Cipher benchmarks
go test -bench=BenchmarkGenerateKeyPair ./ciphering
go test -bench=BenchmarkEncapsulate ./ciphering
go test -bench=BenchmarkDecapsulate ./ciphering

# Signing benchmarks  
go test -bench=BenchmarkSign ./signing
go test -bench=BenchmarkVerify ./signing

# Hashing benchmarks
go test -bench=BenchmarkHash ./hashing
go test -bench=BenchmarkHashSizes ./hashing
```

### Performance Profiling
```bash
# CPU profiling
go test -bench=BenchmarkEncapsulate -cpuprofile=cpu.prof ./ciphering
go tool pprof cpu.prof

# Memory profiling
go test -bench=BenchmarkSign -memprofile=mem.prof ./signing  
go tool pprof mem.prof
```

## Advanced Test Options

### Test with Race Detection
```bash
# Check for race conditions
go test -race ./...
```

### Test with Build Tags
```bash
# Run tests with specific build tags
go test -tags=debug ./...
```

### Timeout and Parallel Execution
```bash
# Set test timeout
go test -timeout=30s ./...

# Control parallel execution
go test -parallel=4 ./...
```

## Coverage Analysis

### Generate Coverage Reports
```bash
# Basic coverage
go test -cover ./...

# Detailed coverage by package
go test -coverprofile=util.out ./util
go test -coverprofile=cipher.out ./ciphering  
go test -coverprofile=signing.out ./signing
go test -coverprofile=hashing.out ./hashing

# Merge coverage reports
go run golang.org/x/tools/cmd/cover merge util.out cipher.out signing.out hashing.out -o total.out

# Generate HTML report
go tool cover -html=total.out -o coverage.html
```

### Coverage by Function
```bash
# Show coverage by function
go tool cover -func=coverage.out
```

## Test Output Examples

### Successful Test Run
```
$ go test -v ./cipher
=== RUN   TestGenerateKeyPair
=== RUN   TestGenerateKeyPair/Level_1_(~AES-128)
=== RUN   TestGenerateKeyPair/Level_3_(~AES-192)
=== RUN   TestGenerateKeyPair/Level_5_(~AES-256)
--- PASS: TestGenerateKeyPair (0.02s)
    --- PASS: TestGenerateKeyPair/Level_1_(~AES-128) (0.01s)
    --- PASS: TestGenerateKeyPair/Level_3_(~AES-192) (0.01s)
    --- PASS: TestGenerateKeyPair/Level_5_(~AES-256) (0.01s)
=== RUN   TestEncapsulateDecapsulate
=== RUN   TestEncapsulateDecapsulate/Level_1_(~AES-128)
=== RUN   TestEncapsulateDecapsulate/Level_3_(~AES-192)
=== RUN   TestEncapsulateDecapsulate/Level_5_(~AES-256)
--- PASS: TestEncapsulateDecapsulate (0.03s)
    --- PASS: TestEncapsulateDecapsulate/Level_1_(~AES-128) (0.01s)
    --- PASS: TestEncapsulateDecapsulate/Level_3_(~AES-192) (0.01s)
    --- PASS: TestEncapsulateDecapsulate/Level_5_(~AES-256) (0.01s)
PASS
ok      trial_pqc/cipher        0.064s
```

### Benchmark Output
```
$ go test -bench=. -benchmem ./cipher
BenchmarkGenerateKeyPair/Kyber512-8       1000      1534 ns/op    2560 B/op      5 allocs/op
BenchmarkGenerateKeyPair/Kyber768-8        800      1876 ns/op    3584 B/op      5 allocs/op  
BenchmarkGenerateKeyPair/Kyber1024-8       600      2234 ns/op    4736 B/op      5 allocs/op
BenchmarkEncapsulate-8                     2000       867 ns/op    1120 B/op      3 allocs/op
BenchmarkDecapsulate-8                     1500       976 ns/op      64 B/op      2 allocs/op
PASS
ok      trial_pqc/cipher        8.234s
```

## Continuous Integration

### GitHub Actions Example
```yaml
name: Tests
on: [push, pull_request]
jobs:
  test:
    runs-on: ubuntu-latest
    steps:
    - uses: actions/checkout@v3
    - uses: actions/setup-go@v3
      with:
        go-version: 1.21
    - run: go test -race -coverprofile=coverage.out ./...
    - run: go tool cover -html=coverage.out -o coverage.html
```

## Makefile for Common Tasks

Create a `Makefile` for convenience:
```makefile
.PHONY: test bench cover clean

test:
	go test -v ./...

test-race:
	go test -race ./...

bench:
	go test -bench=. -benchmem ./...

cover:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

clean:
	rm -f coverage.out coverage.html *.prof
```

Then run:
```bash
make test      # Run all tests
make bench     # Run benchmarks  
make cover     # Generate coverage report
make clean     # Clean up generated files
```