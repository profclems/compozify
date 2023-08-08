package api

import (
	"bytes"
	"context"
	"fmt"
	"net"
	"testing"
	"time"

	"github.com/rs/zerolog"
)

func TestServer(t *testing.T) {
	var logBuffer bytes.Buffer
	logger := zerolog.New(&logBuffer).With().Timestamp().Logger()
	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		t.Fatalf("Failed to create listener: %v", err)
	}
	defer func() {
		if err := listener.Close(); err != nil {
			logger.Error().Err(err).Msg("Failed to close listener")
		}
	}()

	addr := listener.Addr().(*net.TCPAddr)
	port := addr.Port
	endpoint := fmt.Sprintf("http://localhost:%d/api/parse", port)
	logger.Info().Msgf("Endpoint: %s", endpoint)

	server := NewServer(&logger, listener, nil)

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	go func() {
		if err := server.Run(ctx); err != nil && err != context.Canceled {
			logger.Error().Err(err).Msg("Server error")
		}
	}()

	// Wait for the server to start
	time.Sleep(1 * time.Second)

	fmt.Printf("Server is running on %s. You have 2 minutes to test manually.\n", endpoint)

	// Wait for 2 minutes to allow manual testing
	time.Sleep(2 * time.Minute)

}
