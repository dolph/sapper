package main

import (
	"net/http"

	// Request routing
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	r := mux.NewRouter()
	r.NotFoundHandler = http.HandlerFunc(notFound)
	r.HandleFunc("/", echoRemoteAddr).Methods("GET")
	return r
}
