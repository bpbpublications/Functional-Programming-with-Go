package main

import (
    "fmt"
    "log"
    "net/http"
)

// LoggerMiddleware logs the details of incoming requests.
func LoggerMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        log.Printf("Request received - Method: %s, Path: %s\n", r.Method, r.URL.Path)
        next.ServeHTTP(w, r) // Call the next handler
    })
}

// AuthenticationMiddleware checks if the request is authorized.
func AuthenticationMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        apiKey := r.Header.Get("X-API-Key")
        if apiKey != "secret-key" {
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return
        }
        next.ServeHTTP(w, r) // Call the next handler if authorized
    })
}

// HomeHandler is our main HTTP handler.
func HomeHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Welcome to the secure area!")
}

func main() {
    // Wrap the HomeHandler with our middleware
    pipeline := LoggerMiddleware(AuthenticationMiddleware(http.HandlerFunc(HomeHandler)))

    http.Handle("/", pipeline)
    
    log.Println("Server started on :8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
}
