ch := make(chan int, 2)
package main

import "fmt"

func main() {
    messages := make(chan string, 2)

    messages <- "hello"
    messages <- "world"

    fmt.Println(<-messages)
    fmt.Println(<-messages)
}
