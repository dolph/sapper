package main

import (
    "fmt"
    "net/http"
    "strings"
    "google.golang.org/appengine"
)

// Trivially routes requests to the correct response function.
func IndexHandler(w http.ResponseWriter, r *http.Request) {
    switch r.URL.Path {
        case "/":
            switch r.Method {
                case http.MethodGet:
                    GetIndex(w, r)
                default:
                    NotFound(w, r)
            }
        default:
            NotFound(w, r)
    }
}

// Returns the client's IP address.
func GetIndex(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    w.Header().Set("Content-Type", "text/plain; charset=UTF-8")

    // RemoteAddr is formatted as host:port, so we just trim off the port here
    // and return the IP.
    fmt.Fprintf(w, "%s\n", strings.Split(r.RemoteAddr, ":")[0])
}

// Returns a 404 Not Found page.
func NotFound(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusNotFound)
    w.Header().Set("Content-Type", "text/plain; charset=UTF-8")
    fmt.Fprintf(w, "%d Not Found\n", http.StatusNotFound)
}

// appengine.Main() expects packages to register HTTP handlers in their init()
// functions.
func init() {
    http.HandleFunc("/", IndexHandler)
}

func main() {
    // Starts listening on port 8080 (or $PORT), and never returns.
    // https://godoc.org/google.golang.org/appengine#Main
    appengine.Main()
}
