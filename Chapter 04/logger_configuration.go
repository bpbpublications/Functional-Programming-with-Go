func ConfigureLogger(env string) func(string) {
    return func(message string) {
        if env == "production" {
            // Production logging
        } else {
            // Development or staging logging
            fmt.Println(message)
        }
    }
}