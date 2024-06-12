package main

import (
 "fmt"
 "sync"
)

// processItem is a higher-order function that takes a func(int) int and returns a new function.
// This new function is designed to process items concurrently using goroutines.
func processItem(worker func(int) int) func([]int) []int {
 return func(items []int) []int {
  var wg sync.WaitGroup
  results := make([]int, len(items))
  for i, item := range items {
   wg.Add(1)
   // Start a goroutine for each item.
   go func(i, val int) {
    defer wg.Done()
    // Process the item with the provided worker function.
    results[i] = worker(val)
   }(i, item)
  }
  wg.Wait() // Wait for all goroutines to complete.
  return results
 }
}

func main() {
 // Example worker function that squares numbers.
 square := func(num int) int {
  return num * num
 }

 // Create a concurrent processor using the higher-order function.
 concurrentProcessor := processItem(square)

 // Example data slice.
 data := []int{1, 2, 3, 4, 5}

 // Process data concurrently.
 result := concurrentProcessor(data)
 fmt.Println("Processed data:", result)
}
