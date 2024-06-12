package main

import "fmt"

func Map(slice []int, fn func(int) int) []int {
 var result []int
 for _, v := range slice {
  result = append(result, fn(v))
 }
 return result
}

func main() {
 numbers := []int{1, 2, 3, 4}
 squared := Map(numbers, func(n int) int { return n * n })
 fmt.Println("squared: ", squared) // squared: [1, 4, 9, 16]
}
