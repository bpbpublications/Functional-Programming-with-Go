package main

import "fmt"

// Counter returns a tuple of functions: increment, decrement, and getValue to manage the state of the counter
func Counter(initial int) (func(), func(), func() int) {
    value := initial // value is enclosed by the returned functions

    // increment adds one to the counter
    increment := func() {
        value++
    }

    // decrement subtracts one from the counter
    decrement := func() {
        value--
    }

    // getValue returns the current value of the counter
    getValue := func() int {
        return value
    }

    return increment, decrement, getValue
}

func main() {
    // Initialize the counter with a starting value of 10
    inc, dec, getVal := Counter(10)

    fmt.Println("Initial Value:", getVal()) // Outputs: Initial Value: 10

    inc() // Increment the counter
    fmt.Println("After incrementing:", getVal()) // Outputs: After incrementing: 11

    dec() // Decrement the counter
    dec() // Decrement the counter again
    fmt.Println("After decrementing twice:", getVal()) // Outputs: After decrementing twice: 9
}
