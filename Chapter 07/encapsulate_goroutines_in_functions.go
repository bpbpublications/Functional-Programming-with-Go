package main

import (
    "fmt"
)

// processData processes data and returns a result
func processData(data int) int {
    // Simulate data processing
    return data * 2
}

// startGoroutine wraps a data processing function in a goroutine
func startGoroutine(data int, resultChan chan<- int) {
    result := processData(data)
    resultChan <- result
}

func main() {
    data := []int{1, 2, 3, 4, 5}
    resultChan := make(chan int, len(data))

    for _, v := range data {
        go startGoroutine(v, resultChan)
    }

    for range data {
        fmt.Println(<-resultChan)
    }
}
