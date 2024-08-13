# Define variables
APP_NAME := go-starter-template

# Default target
all: build

# Build the Go application
build:
	@echo "Building the application..."
	go build -o $(APP_NAME) ./cmd/todo-service

# Clean the build
clean:
	@echo "Cleaning up..."
	rm -f $(APP_NAME)

# Run the application
run: build
	@echo "Running the application..."
	./$(APP_NAME)

# Generate protobuf files
proto:
	@echo "Generating protobuf files..."
	protoc -I. --go_out=paths=source_relative:. --go-grpc_out=paths=source_relative:./ rpc/todo/todo.proto

# Format the code
fmt:
	@echo "Formatting code..."
	go fmt ./...

# Lint the code
lint:
	@echo "Linting code..."
	golangci-lint run

.PHONY: all build clean run proto test fmt lint
