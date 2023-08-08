package api

import (
	"encoding/json"
	"fmt"
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

	start := time.Now()
	code := http.StatusOK
	errorMsg := ""
	defer func() {
		log := logger.Info()
		if errorMsg != "" {
			log = logger.Error()
			http.Error(w, errorMsg, code)
		}
		log.Msgf("Returned %d in %v", code, time.Since(start))
	}()

	err := json.NewDecoder(r.Body).Decode(&dockerCmds)
	if err != nil {
		errorMsg = fmt.Sprintf("Error decoding request body: %v", err)
		code = http.StatusBadRequest
		return
	}

	var p *parser.Parser
	for _, cmd := range dockerCmds.Commands {
		if cmd == "" {
			errorMsg = "Docker command cannot be empty"
			code = http.StatusBadRequest
			return
		}

		// Create a new Parser or append to existing parser
		if p == nil {
			p, err = parser.New(cmd)
		} else {
			yamlBytes := []byte(p.String())
			p, err = parser.AppendToYAML(yamlBytes, cmd)
		}

		if err != nil {
			errorMsg = fmt.Sprintf("Error parsing Docker command: %v", err)
			code = http.StatusBadRequest
			return
		}

		err = p.Parse()
		if err != nil {
			errorMsg = fmt.Sprintf("Error parsing Docker command: %v", err)
			code = http.StatusBadRequest
			return
		}
	}

	dockerComposeYaml := p.String()

	// Create the response
	resp := Response{
		Output: dockerComposeYaml,
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		logger.Err(err).Msg("Unable to write response")
		return
	}
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
