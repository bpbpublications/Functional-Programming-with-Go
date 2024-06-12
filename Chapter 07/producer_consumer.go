package main

import (
    "fmt"
    "time"
)

// producer generates data and sends it to a channel
func producer(ch chan<- int) {
    for i := 0; i < 5; i++ {
        ch <- i
        time.Sleep(time.Second)
    }
    close(ch)
}

// consumer reads data from a channel and processes it
func consumer(ch <-chan int) {
    for data := range ch {
        fmt.Println("Processed:", processData(data))
    }
}

func main() {
    ch := make(chan int)
    go producer(ch)
    consumer(ch)
}
