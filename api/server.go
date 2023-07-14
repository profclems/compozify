package api

import (
	"context"
	"errors"
	"io"
	"io/fs"
	"net"
	"net/http"
	"path"
	"strings"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
	"golang.org/x/sync/errgroup"

	"github.com/profclems/compozify/api/router"
)

type Server struct {
	logger   *zerolog.Logger
	listener net.Listener
	http     http.Server
	assets   fs.FS
}

func NewServer(logger *zerolog.Logger, listener net.Listener, assets fs.FS) *Server {
	server := &Server{
		logger:   logger,
		listener: listener,
		assets:   assets,
	}
	r := mux.NewRouter()
	router.Handle(r)

	r.PathPrefix("/static").HandlerFunc(server.cacheHandler)
	r.PathPrefix("/").HandlerFunc(server.appHandler)

	server.http = http.Server{
		Handler: r,
	}

	return server
}

// appHandler is web app http handler function.
func (server *Server) appHandler(w http.ResponseWriter, r *http.Request) {
	header := w.Header()

	header.Set("Content-Type", "text/html; charset=UTF-8")
	header.Set("X-Content-Type-Options", "nosniff")
	header.Set("Referrer-Policy", "same-origin")

	f, err := server.assets.Open("index.html")
	if err != nil {
		http.Error(w, `web/ unbuilt`, http.StatusNotFound)
		return
	}
	defer func() { _ = f.Close() }()

	_, _ = io.Copy(w, f)
}

func (server *Server) cacheHandler(w http.ResponseWriter, r *http.Request) {
	staticServer := http.FileServer(http.FS(server.assets))
	header := w.Header()

	if contentType, ok := commonContentType(path.Ext(r.URL.Path)); ok {
		header.Set("Content-Type", contentType)
	}

	header.Set("Cache-Control", "public, max-age=31536000")
	header.Set("X-Content-Type-Options", "nosniff")
	header.Set("Referrer-Policy", "same-origin")

	staticServer.ServeHTTP(w, r)
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

func commonContentType(ext string) (string, bool) {
	ext = strings.ToLower(ext)
	mime, ok := commonTypes[ext]
	return mime, ok
}

var commonTypes = map[string]string{
	".css":   "text/css; charset=utf-8",
	".gif":   "image/gif",
	".htm":   "text/html; charset=utf-8",
	".html":  "text/html; charset=utf-8",
	".jpeg":  "image/jpeg",
	".jpg":   "image/jpeg",
	".js":    "application/javascript",
	".mjs":   "application/javascript",
	".otf":   "font/otf",
	".pdf":   "application/pdf",
	".png":   "image/png",
	".svg":   "image/svg+xml",
	".ttf":   "font/ttf",
	".wasm":  "application/wasm",
	".webp":  "image/webp",
	".xml":   "text/xml; charset=utf-8",
	".sfnt":  "font/sfnt",
	".woff":  "font/woff",
	".woff2": "font/woff2",
}
