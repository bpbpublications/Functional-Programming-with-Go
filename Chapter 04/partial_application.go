package main

import (
 "fmt"
)

// Function to filter data
func FilterData(criteria func(int) bool, data []int) []int {
 var result []int
 for _, value := range data {
  if criteria(value) {
   result = append(result, value)
  }
 }
 return result
}

// Function to transform data
func TransformData(transformer func(int) int, data []int) []int {
 var result []int
 for _, value := range data {
  result = append(result, transformer(value))
 }
 return result
}

// Function to aggregate data
func AggregateData(aggregator func([]int) int, data []int) int {
 return aggregator(data)
}

// Partially applied filter
// filters out even numbers
func IsEven(n int) bool {
 return n%2 == 0
}

func FilterEvenNumbers(data []int) []int {
 return FilterData(IsEven, data)
}

// Partially applied transformer
// squares the numbers
func Square(n int) int {
 return n * n
}

func SquareNumbers(data []int) []int {
 return TransformData(Square, data)
}

// Partially applied aggregator
// performs sum
func Sum(numbers []int) int {
 total := 0
 for _, number := range numbers {
  total += number
 }
 return total
}

func SumNumbers(data []int) int {
 return AggregateData(Sum, data)
}

func main() {
 data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

 // Apply the data processing pipeline
 filtered := FilterEvenNumbers(data) // Keep even numbers
 transformed := SquareNumbers(filtered) // Square the remaining numbers
 aggregated := SumNumbers(transformed)  // Sum the squared numbers

 fmt.Println("Sum of squares of odd numbers from 1 to 10:", aggregated)
 // Output: Sum of squares of odd numbers from 1 to 10: 220
}
