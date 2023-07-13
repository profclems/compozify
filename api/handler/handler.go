package handler

import (
	"encoding/json"
	"fmt"
	"github.com/profclems/compozify/pkg/parser"
	"net/http"
)

type DockerCommand struct {
	Command string `json:"command"`
}

func ParseDockerCommand(w http.ResponseWriter, r *http.Request) {
	var dockerCmd DockerCommand
	err := json.NewDecoder(r.Body).Decode(&dockerCmd)

	if err != nil {
		http.Error(w, "Error reading request body",
			http.StatusInternalServerError)
		return
	}

	dockerCommand := dockerCmd.Command
	// Create a new Parser
	p, err := parser.NewParser(dockerCommand)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Parse the Docker command
	err = p.Parse()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	dockerComposeYaml := p.String()

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"compose": "%s"}`, dockerComposeYaml)
}
