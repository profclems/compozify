package main

import (
	"github.com/gorilla/mux"
	"github.com/profclems/compozify/api/router"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	router.Handle(r)

	log.Fatal(http.ListenAndServe(":8080", r))
}
