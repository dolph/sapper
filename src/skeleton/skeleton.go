package skeleton

import (
    "fmt"
    "net/http"
    "google.golang.org/appengine"
)

func init() {
    http.HandleFunc("/", IndexHandler)
}

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

func GetIndex(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    w.Header().Set("Content-Type", "text/plain")
    fmt.Fprint(w, "hello, world")
}

func NotFound(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusNotFound)
    w.Header().Set("Content-Type", "text/plain")
    fmt.Fprintf(w, "%d Not Found", http.StatusNotFound)
}

func main() {
    appengine.Main()
}
