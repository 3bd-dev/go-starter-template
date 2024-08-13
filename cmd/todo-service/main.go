package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/3bd-dev/go-starter-template/internal/handlers/todoapi"
	"github.com/3bd-dev/go-starter-template/internal/handlers/todogrpc"
	"github.com/3bd-dev/go-starter-template/internal/repos/inmemory"
	"github.com/3bd-dev/go-starter-template/internal/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// Repository setup (choose between in-memory or PostgreSQL)
	repo := inmemory.NewTodoRepository()
	// repo := postgres.NewTodoRepository(db)

	// Service setup
	todoService := services.NewTodoService(repo)

	// Channel for OS signals
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// grpc server
	grpcServer := grpc.NewServer()
	todogrpc.Register(grpcServer, todoService)
	reflection.Register(grpcServer)

	// http server
	httpmux := http.NewServeMux()
	todoapi.Routes(httpmux, todoService)

	httpserver := &http.Server{
		Addr:    ":8080",
		Handler: httpmux,
	}

	var wg sync.WaitGroup

	// Start gRPC server
	wg.Add(1)
	go func() {
		defer wg.Done()
		lis, err := net.Listen("tcp", ":8003")
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}

		log.Println("Starting gRPC server on :8003...")
		if err := grpcServer.Serve(lis); err != nil && err != grpc.ErrServerStopped {
			log.Fatalf("failed to serve: %v", err)
		}
		log.Println("gRPC server stopped")
	}()

	// Start HTTP server
	wg.Add(1)
	go func() {
		defer wg.Done()

		go func() {
			<-quit
			log.Println("Shutting down HTTP server...")
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			if err := httpserver.Shutdown(ctx); err != nil {
				log.Printf("HTTP server Shutdown: %v", err)
			}
			log.Println("HTTP server stopped")
		}()

		log.Println("Starting HTTP server on :8080...")
		if err := httpserver.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	shutdown := <-quit
	log.Println("Received shutdown signal")

	// Shut down the gRPC server
	log.Println("Initiating gRPC server shutdown...")
	grpcServer.GracefulStop()
	quit <- shutdown
	// Close the quit channel and wait for servers to shut down
	wg.Wait()
	log.Println("Servers shut down gracefully")
}
