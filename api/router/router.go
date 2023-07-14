package router

import (
	"github.com/gorilla/mux"
	"github.com/profclems/compozify/api/handler"
)

func Handle(r *mux.Router) {
	r.HandleFunc("/api/parse", handler.ParseDockerCommand).Methods("POST")
}
