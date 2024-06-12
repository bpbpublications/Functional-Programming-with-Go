package main

import "fmt"

// LazyInts returns a function that generates integers starting from i
func LazyInts(i int) func() int {
    return func() int {
        result := i
        i++
        return result
    }
}

func main() {
    nextInt := LazyInts(0) // Initialize starting at 0

    fmt.Println(nextInt()) // 0
    fmt.Println(nextInt()) // 1
    fmt.Println(nextInt()) // 2
    // Only computes the next integer when called
}
