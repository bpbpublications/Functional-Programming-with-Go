package main

import (
 "fmt"
 "log"
)

func divide(a, b float64) (float64, error) {
 if b == 0 {
  return 0, fmt.Errorf("cannot divide by zero")
 }
 return a / b, nil
}

func main() {
 result, err := divide(10, 0)
 if err != nil {
  log.Fatal(err)
 }
 fmt.Println("Result:", result)

}
