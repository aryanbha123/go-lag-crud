.PHONY: run build tidy test clean

# Run the API server in development
run:
	go run ./cmd/api

# Compile a binary into bin/api
build:
	go build -o bin/api ./cmd/api

# Add missing and remove unused module dependencies
tidy:
	go mod tidy

# Run all tests in the module
test:
	go test ./...

# Remove build artifacts
clean:
	rm -rf bin
