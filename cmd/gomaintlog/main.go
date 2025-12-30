package main

import (
	"context"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/maintenance_log"
	"github.com/mark3labs/mcp-go/server"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	maintlogServer, err := maintenance_log.NewServer()
	if err != nil {
		log.Fatalf("Failed to create maintenance log server: %v", err)
	}

	httpServer := server.NewStreamableHTTPServer(
		maintlogServer.MCPServer(),
		server.WithStateful(true),
	)

	addr := ":8080"
	log.Printf("Starting maintenance log MCP server on %s", addr)
	log.Printf("MCP endpoint available at: http://localhost%s/mcp", addr)

	go func() {
		if err := httpServer.Start(addr); err != nil {
			log.Printf("Server error: %v", err)
		}
	}()

	// Wait for interrupt signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := httpServer.Shutdown(ctx); err != nil {
		fmt.Fprintf(os.Stderr, "Server shutdown error: %v\n", err)
		os.Exit(1)
	}

	log.Println("Server stopped")
}
