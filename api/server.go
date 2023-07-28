package api

import (
	"context"
	"errors"
	"io/fs"
	"net"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
	"golang.org/x/sync/errgroup"
)

// Server is the web app server.
type Server struct {
	logger   *zerolog.Logger
	listener net.Listener
	http     http.Server
	assets   fs.FS
}

// NewServer creates a new Server.
func NewServer(logger *zerolog.Logger, listener net.Listener, assets fs.FS) *Server {
	server := &Server{
		logger:   logger,
		listener: listener,
		assets:   assets,
	}

	r := mux.NewRouter()
	r.HandleFunc("/api/parse", server.ParseDockerCommand).Methods("POST")

	r.PathPrefix("/static").HandlerFunc(server.cacheHandler)
	r.PathPrefix("/").HandlerFunc(server.appHandler)

	server.http = http.Server{
		Handler: r,
	}

	return server
}

// Run starts the server that host webapp and api endpoints.
func (server *Server) Run(ctx context.Context) (err error) {
	ctx, cancel := context.WithCancel(ctx)
	var group errgroup.Group

	group.Go(func() error {
		<-ctx.Done()
		return server.http.Shutdown(context.Background())
	})
	group.Go(func() error {
		defer cancel()
		err := server.http.Serve(server.listener)
		if err == context.Canceled || errors.Is(err, http.ErrServerClosed) {
			err = nil
		}
		return err
	})

	return group.Wait()
}

// Close closes server and underlying listener.
func (server *Server) Close() error {
	return server.http.Close()
}
