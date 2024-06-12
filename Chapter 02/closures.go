package main

import "fmt"

func main() {
    // Outer function
    outer := func() {
        x := 10

        // Inner closure
        inner := func() {
            fmt.Println("Value of x from closure:", x)
        }

        // Invoke the inner closure
        inner()
    }

    // Invoke the outer function
    outer()
}
