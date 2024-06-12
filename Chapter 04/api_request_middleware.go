func LoggingMiddleware(level string) func(http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            // Log request based on the level
            log.Printf("[%s] %s", level, r.URL.Path)
            next.ServeHTTP(w, r)
        })
    }
}
