package main

import (
 "fmt"
 "log"
)

type Result struct {
 Value float64
 Err   error
}

func divideFunc(a, b float64) Result {
 if b == 0 {
  return Result{0, fmt.Errorf("cannot divide by zero")}
 }
 return Result{a / b, nil}
}

func main() {
 result := divideFunc(10, 0)
 if result.Err != nil {
  log.Fatal(result.Err)
 }
 fmt.Println("Result:", result.Value)
}
