package main

import (
 "fmt"
 "sort"
)

// Comparator function type definition
type Comparator func(a, b int) bool

// sortInts sorts the slice of integers using the provided comparator callback
func sortInts(data []int, comparator Comparator) {
 sort.Slice(data, func(i, j int) bool {
  return comparator(data[i], data[j])
 })
}

func main() {
 // Slice of integers to sort
 numbers := []int{9, 5, 7, 3, 1, 6, 4, 8, 2, 0}

 // Define an ascending order comparator
 ascending := func(a, b int) bool {
  return a < b
 }

 // Define a descending order comparator
 descending := func(a, b int) bool {
  return a > b
 }

 // Sort the numbers in ascending order using the ascending comparator
 fmt.Println("Original:", numbers)
 sortInts(numbers, ascending)
 fmt.Println("Sorted in ascending order:", numbers)

 // Sort the numbers in descending order using the descending comparator
 sortInts(numbers, descending)
 fmt.Println("Sorted in descending order:", numbers)
}
