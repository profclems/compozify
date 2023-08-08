package api

import (
	"context"
	"net"
	"os"
	"testing"
	"time"

	"github.com/rs/zerolog"
)

func TestServer(t *testing.T) {
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		t.Fatalf("Failed to create listener: %v", err)
	}

	server := NewServer(&logger, listener, nil)

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	if err := server.Run(ctx); err != nil && err != context.Canceled && err != context.DeadlineExceeded {
		t.Fatalf("Server error: %v", err)
	}
}
