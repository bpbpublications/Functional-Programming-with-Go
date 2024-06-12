package main

import "fmt"

// fib calculates the n-th Fibonacci number
func fib(n int) int {
 if n <= 1 {
  return n
 }
 return fib(n-1) + fib(n-2)
}

func main() {
 fmt.Println(fib(7))  // Output: 13
 fmt.Println(fib(10)) // Output: 55
}
