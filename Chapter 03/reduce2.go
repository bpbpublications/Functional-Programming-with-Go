package main

import "fmt"

func Reduce(slice []int, initializer int, op func(int, int) int) int {
 acc := initializer
 for _, v := range slice {
  acc = op(acc, v)
 }
 return acc
}

func main() {
 numbers := []int{1, 2, 3, 4}
 sum := Reduce(numbers, 0, func(a, b int) int { return a + b })
 fmt.Println("sum: ", sum) // // sum: 10
}
