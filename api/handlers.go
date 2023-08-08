package api

import (
	"encoding/json"
	"fmt"
	"github.com/rs/zerolog"
	"io"
	"net/http"
	"path"
	"strings"
	"time"

	"github.com/profclems/compozify/pkg/parser"
)

// Response is the response body for the ParseDockerCommand handler.
type Response struct {
	Output string `json:"output"`
}

// ParseDockerCommands ParseDockerCommand parses a Docker command and returns the equivalent Docker Compose YAML.
func (server *Server) ParseDockerCommands(w http.ResponseWriter, r *http.Request) {
	type DockerCommands struct {
		Commands []string `json:"commands"`
	}
	var dockerCmds DockerCommands

	logger := server.logger.With().Str("handler", "ParseDockerCommands").Str("remoteAddr", r.RemoteAddr).Logger()
	logger.Info().Msgf("%s %s %s", r.Method, r.URL.Path, r.Proto)

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			logger.Err(err).Msg("Error closing request body")
		}
	}(r.Body)

	err := json.NewDecoder(r.Body).Decode(&dockerCmds)
	if err != nil {
		writeError(w, logger, "Error decoding request body", err, http.StatusBadRequest)
		return
	}

	start := time.Now()
	defer logDuration(logger, start)

	var p *parser.Parser
	for _, cmd := range dockerCmds.Commands {
		if cmd == "" {
			writeError(w, logger, "Docker command cannot be empty", nil, http.StatusBadRequest)
			return
		}

		if p == nil {
			p, err = parser.New(cmd)
		} else {
			p, err = parser.AppendToYAML([]byte(p.String()), cmd)
		}

		if err != nil {
			writeError(w, logger, "Error parsing Docker command", err, http.StatusBadRequest)
			return
		}

		if err := p.Parse(); err != nil {
			writeError(w, logger, "Error parsing Docker command", err, http.StatusBadRequest)
			return
		}
	}

	resp := Response{
		Output: p.String(),
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		logger.Err(err).Msg("Unable to write response")
	}
}

func writeError(w http.ResponseWriter, logger zerolog.Logger, msg string, err error, code int) {
	logger.Err(err).Msg(msg)
	http.Error(w, fmt.Sprintf("%s: %v", msg, err), code)
}

func logDuration(logger zerolog.Logger, start time.Time) {
	logger.Info().Msgf("Returned in %v", time.Since(start))
}

// appHandler is web app http handler function.
func (server *Server) appHandler(w http.ResponseWriter, r *http.Request) {
	staticServer := http.FileServer(http.FS(server.assets))
	header := w.Header()

	if contentType, ok := commonContentType(path.Ext(r.URL.Path)); ok {
		header.Set("Content-Type", contentType)
	}

	header.Set("X-Content-Type-Options", "nosniff")
	header.Set("Referrer-Policy", "same-origin")

	staticServer.ServeHTTP(w, r)
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
