package main

import (
 "fmt"
 "time"
)

func main() {
 // Creating an unbuffered channel
 unbufferedChannel := make(chan string)

 // Creating a buffered channel with a capacity of 2
 bufferedChannel := make(chan string, 2)

 // Goroutine for sending data to the unbuffered channel
 go func() {
  fmt.Println("Sending data to unbuffered channel...")
  unbufferedChannel <- "Hello from unbuffered channel!"
  fmt.Println("Data sent to unbuffered channel")
 }()

 // Goroutine for receiving data from the unbuffered channel
 go func() {
  // Simulating delay
  time.Sleep(1 * time.Second)
  data := <-unbufferedChannel
  fmt.Println("Received:", data)
 }()

 // Sending data to the buffered channel
 fmt.Println("Sending data to buffered channel...")
 bufferedChannel <- "First message"
 bufferedChannel <- "Second message"
 fmt.Println("Data sent to buffered channel without blocking")

 // Goroutine for receiving data from the buffered channel
 go func() {
  // Simulating delay
  time.Sleep(2 * time.Second)
  for i := 0; i < 2; i++ {
   data := <-bufferedChannel
   fmt.Println("Received from buffered:", data)
  }
 }()

 // Wait for a key press to exit
 fmt.Println("Press Enter to exit...")
 fmt.Scanln()
 fmt.Println("Exiting...")
}
