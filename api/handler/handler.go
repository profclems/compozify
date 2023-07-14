package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/profclems/compozify/pkg/parser"
)

type DockerCommand struct {
	Command string `json:"command"`
}

func ParseDockerCommand(w http.ResponseWriter, r *http.Request) {
	var dockerCmd DockerCommand
	err := json.NewDecoder(r.Body).Decode(&dockerCmd)

	if err != nil {
		http.Error(w, fmt.Sprintf("Error decoding request body: %v", err), http.StatusBadRequest)
		return
	}

	// Validate the command.
	if dockerCmd.Command == "" {
		http.Error(w, "Docker command cannot be empty", http.StatusBadRequest)
		return
	}

	// Create a new Parser
	p, err := parser.NewParser(dockerCmd.Command)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error creating parser: %v", err), http.StatusBadRequest)
		return
	}

	// Parse the Docker command
	err = p.Parse()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error parsing command: %v", err), http.StatusBadRequest)
		return
	}

	dockerComposeYaml := p.String()

	w.Header().Set("Content-Type", "application/x-yaml")
	_, err = w.Write([]byte(dockerComposeYaml))
	if err != nil {
		log.Printf("Unable to write response: %v", err)
		return
	}
}
