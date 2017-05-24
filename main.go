package main

import (
	"net/http"

	// Google App Engine
	"google.golang.org/appengine"
)

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
