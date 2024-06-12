package main

import "fmt"

// fibMemo wraps a memoized Fibonacci function
func fibMemo() func(int) int {
    memo := make(map[int]int)
    var fib func(int) int

    fib = func(n int) int {
        if n <= 1 {
            return n
        }
        // Check if the result is already in the memo map
        if result, found := memo[n]; found {
            return result
        }
        // Store the result in the memo map
        result := fib(n-1) + fib(n-2)
        memo[n] = result
        return result
    }

    return fib
}

func main() {
    fib := fibMemo()
    fmt.Println("Fibonacci of 10:", fib(10))
    fmt.Println("Fibonacci of 15:", fib(15))
    fmt.Println("Fibonacci of 20:", fib(20))
}
