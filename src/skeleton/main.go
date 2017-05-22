package main

import (
	"fmt"
	"net/http"
	"strings"

	"google.golang.org/appengine"

	// Request routing
	"github.com/gorilla/mux"
)

// Returns the client's IP address.
func GetIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	// RemoteAddr is formatted as host:port, so we just trim off the port here
	// and return the IP.
	var ip string
	switch index := strings.LastIndex(r.RemoteAddr, ":"); index {
	case -1:
		ip = r.RemoteAddr
	default:
		ip = r.RemoteAddr[:index]
	}
	fmt.Fprintf(w, "%s\n", ip)
}

// Returns a 404 Not Found page.
func NotFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, "%d Not Found\n", http.StatusNotFound)
}

func Router() *mux.Router {
	r := mux.NewRouter()
	r.NotFoundHandler = http.HandlerFunc(NotFound)
	r.HandleFunc("/", GetIndex).Methods("GET")
	return r
}

// appengine.Main() expects packages to register HTTP handlers in their init()
// functions.
func init() {
	http.Handle("/", Router())
}

func main() {
	// Starts listening on port 8080 (or $PORT), and never returns.
	// https://godoc.org/google.golang.org/appengine#Main
	appengine.Main()
}
