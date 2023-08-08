package api

import (
	"context"
	"net"
	"os"
	"testing"

	"github.com/rs/zerolog"
)

func TestServer(t *testing.T) {
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		t.Fatalf("Failed to create listener: %v", err)
	}

	server := NewServer(&logger, listener, nil)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err := server.Run(ctx); err != nil {
		t.Fatalf("Server error: %v", err)
	}
}
