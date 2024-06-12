package main

import (
 "fmt"
 "net/http"
)

func loggerMiddleware(next http.HandlerFunc) http.HandlerFunc {
 return func(w http.ResponseWriter, r *http.Request) {
  fmt.Printf("Request received: %s %s\n", r.Method, r.URL)
  next(w, r)
 }
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
 fmt.Fprint(w, "Hello, World!")
}

func main() {
 http.HandleFunc("/", loggerMiddleware(helloHandler))
 http.ListenAndServe(":8080", nil)
}
