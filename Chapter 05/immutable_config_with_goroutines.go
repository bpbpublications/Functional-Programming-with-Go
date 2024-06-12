type ImmutableConfig struct {
    settings map[string]string
}

func (c *ImmutableConfig) Get(key string) string {
    return c.settings[key]
}

func handleRequest(config *ImmutableConfig) {
    // Perform operations using config
}

func main() {
    config := &ImmutableConfig{settings: map[string]string{"mode": "production"}}
    go handleRequest(config)
    go handleRequest(config)
    // The immutable config can be safely shared across goroutines
}
