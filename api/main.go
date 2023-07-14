package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/profclems/compozify/api/router"
)

func main() {
	r := mux.NewRouter()
	router.Handle(r)

	log.Fatal(http.ListenAndServe(":8080", r))
}
