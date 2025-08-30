# Variables
BINARY_NAME=go-wc
INPUT ?= 
OPTIONS ?=

## Build the application
build:
	go build -o bin/$(BINARY_NAME) ./cmd/go-wc

## Run the application
run:
	go run ./cmd/go-wc $(if $(OPTIONS),-$(OPTIONS)) $(INPUT)

## Test with file
test-run: build
	./bin/$(BINARY_NAME) $(if $(OPTIONS),-$(OPTIONS)) $(INPUT)

## Clean build artifacts
clean:
	rm -rf bin/

## Format code
fmt:
	go fmt ./...

## Show help
help:
	@echo "Available targets:"
	@echo "  run       - Run with stdin (or make run FILE=filename)"
	@echo "  test-run  - Run with test.txt"
	@echo "  build     - Build the application"
	@echo "  clean     - Remove build artifacts"
	@echo "  fmt       - Format code"
	@echo "  help      - Show this help"

.PHONY: build run test-run clean fmt help
