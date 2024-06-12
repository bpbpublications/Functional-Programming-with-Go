package main

import (
 "fmt"
)

func GenerateInts() <-chan int {
 ch := make(chan int)
 go func() {
  i := 0
  for {
   ch <- i
   i++
  }
 }()
 return ch
}

func main() {
 ints := GenerateInts()
 fmt.Println(<-ints) // 0
 fmt.Println(<-ints) // 1
 fmt.Println(<-ints) // 2
}
