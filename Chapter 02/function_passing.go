package main

import "fmt"

func calculate(x, y int, op func(int, int) int) int {
 return op(x, y)
}

func main() {
 add := func(a, b int) int {
  return a + b
 }

 subtract := func(a, b int) int {
  return a - b
 }

 result := calculate(5, 3, add)
 fmt.Println("Result:", result)
 result = calculate(5, 3, subtract)
 fmt.Println("Result:", result)
}
