# Define variables
APP_NAME := go-starter-template
PROTO_DIR := rpc
PROTO_FILES := $(PROTO_DIR)/*.proto
GO_OUT := $(PROTO_DIR)

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
	protoc --go_out=$(GO_OUT) --go-grpc_out=$(GO_OUT) $(PROTO_FILES)

# Run tests
test:
	@echo "Running tests..."
	go test ./...

# Format the code
fmt:
	@echo "Formatting code..."
	go fmt ./...

# Lint the code
lint:
	@echo "Linting code..."
	golangci-lint run

.PHONY: all build clean run proto test fmt lint
