package main

import "fmt"

// Declare a function type
type MathFunc func(int, int) int

// A function that takes two integers and a MathFunc as arguments
func calculate(x, y int, op MathFunc) int {
    return op(x, y)
}

func main() {
    // Define a function literal (anonymous function)
    add := func(a, b int) int {
        return a + b
    }

    // Pass the add function as an argument
    result := calculate(5, 3, add)

    fmt.Println("Result of addition:", result)
}
