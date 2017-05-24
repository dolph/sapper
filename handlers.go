package main

import (
	"fmt"
	"net/http"
	"strings"
)

// Returns the client's IP address.
func getIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	// RemoteAddr is formatted as host:port, so we just trim off the port here
	// and return the IP.
	var ip string
	switch strings.Count(r.RemoteAddr, ":") {
	case 1:
		// IPv4 addresses may be of the form IP:port
		index := strings.LastIndex(r.RemoteAddr, ":")
		ip = r.RemoteAddr[:index]
	default:
		// IPv6 addresses have multiple colons, and no ports.
		ip = r.RemoteAddr
	}
	fmt.Fprintf(w, "%s\n", ip)
}

// Returns a 404 Not Found page.
func notFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, "%d Not Found\n", http.StatusNotFound)
}
