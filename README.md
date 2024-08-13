# Go Starter Template

This Go starter template offers a straightforward foundation for building a Go application with both HTTP and gRPC services. It's designed to help you get started on your project and guide your thought process on structuring your service. The key is to "start simple while keeping the future in mind"—avoid jumping into over-engineering right from the beginning.

## Project Structure

```bash
go-starter-template/
├── cmd/
│   └── server/
│       └── main.go                 # Main entry point for the application
├── rpc/
│   ├── todo.proto                  # Protobuf definition file
│   ├── todo.pb.go                  # Generated Go file from the protobuf
│   └── todo_grpc.pb.go             # Generated Go file for gRPC
├── internal/
│   ├── handlers/
│   │   ├── http/
│   │   │   └── todo_handler.go      # HTTP handlers for the Todo service
│   │   ├── grpc/
│   │   │   └── todo_handler.go      # gRPC handlers for the Todo service
│   ├── services/
│   │   └── todo_service.go         # Business logic for the Todo service
│   ├── models/
│   │   └── todo.go                 # Data model for the Todo service
│   ├── repos/
│   │   ├── inmemory/
│   │        └── todo_repository.go   # In-memory repository implementation
├── Makefile                        # Makefile for build automation
├── go.mod                          # Go module file
└── go.sum                          # Go module dependencies
```

## Prerequisites

- `protoc` (Protocol Buffers compiler)

### Installation

1. **Clone the repository**:

   ```bash
   git clone https://github.com/your-username/go-starter-template.git
   cd go-starter-template
   ```

2. **Install dependencies**:

   ```bash
   go mod tidy
   ```

3. **Generate protobuf files**:

   ```bash
   make proto
   ```

4. **Build the application**:

   ```bash
   make build
   ```

5. **Run the application**:

   ```bash
   make run
   ```

### Usage

This template includes both an HTTP server and a gRPC server running concurrently. You can interact with the Todo service via:

- **HTTP**: Default port `8080`.
- **gRPC**: Default port `8003`.


## Note

This repo is just to help you get started and is not ready to use as-is. There are a lot of missing functionalities and packages, such as logging, tracing, testing, and more.