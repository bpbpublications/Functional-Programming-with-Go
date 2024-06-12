package main

import (
    "fmt"
    "net/http"
    "time"
)

// handleRequest simulates processing a network request.
func handleRequest(id int) {
    fmt.Printf("Starting request %d\n", id)
    // Simulate a request handling time.
    time.Sleep(time.Duration(id) * time.Second)
    fmt.Printf("Finished request %d\n", id)
}

func main() {
    // Start several goroutines to handle multiple requests concurrently.
    for i := 1; i <= 5; i++ {
        go handleRequest(i)
    }

    // Wait for some time to allow all goroutines to finish their execution.
    fmt.Println("Waiting for all requests to finish...")
    time.Sleep(10 * time.Second)

    fmt.Println("All requests processed.")
}
