package main

import "fmt"

func Filter(slice []int, predicate func(int) bool) []int {
 var result []int
 for _, v := range slice {
  if predicate(v) {
   result = append(result, v)
  }
 }
 return result
}

func main() {
 numbers := []int{1, 2, 3, 4}
 evens := Filter(numbers, func(n int) bool { return n%2 == 0 })
 fmt.Println("evens: ", evens) // // evens: [2, 4]
}
