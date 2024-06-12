package main

import (
    "fmt"
    "log"
    "net/http"
)

// Logger is a middleware function that logs the request details.
func Logger(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        log.Printf("Received request: %s %s", r.Method, r.URL.Path)
        next(w, r)
    }
}

// Authenticate is a middleware function that simulates request authentication.
func Authenticate(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        apiKey := r.Header.Get("X-API-KEY")
        if apiKey != "secret" {
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return
        }
        next(w, r)
    }
}

// HomeHandler is the final handler that responds to HTTP requests.
func HomeHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Welcome to the secure section!")
}

func main() {
    // Wrap the HomeHandler with Logger and Authenticate middleware.
    http.Handle("/", Logger(Authenticate(HomeHandler)))
    log.Println("Server starting on port 8080...")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatalf("Error starting server: %s\n", err)
    }
}
