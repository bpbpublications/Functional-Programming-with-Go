package main

import (
    "fmt"
    "time"
)

func worker(done chan bool) {
    fmt.Print("Working...")
    time.Sleep(time.Second)
    fmt.Println("done")

    // Send a value to notify that we're done.
    done <- true
}

func main() {
    done := make(chan bool)
    go worker(done)

    // Block until we receive a notification from the worker
    <-done
}
