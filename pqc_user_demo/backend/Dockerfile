FROM golang:1.22-alpine AS builder

WORKDIR /app

# Copy go module files first
COPY go.mod go.sum ./

# Copy vendor directory (pre-vendored: run `go mod vendor` or use existing)
COPY vendor/ vendor/

# Copy all Go source files
COPY *.go ./

# Build with vendor flag (offline — no network needed)
RUN GOFLAGS="-mod=vendor" go build -ldflags="-s -w" -o pqc-api .

# ── Runtime image ─────────────────────────────────────────────────────────────
FROM alpine:3.19

RUN apk add --no-cache ca-certificates tzdata

WORKDIR /app
COPY --from=builder /app/pqc-api .

EXPOSE 8080

ENV PORT=8080

ENTRYPOINT ["./pqc-api"]
