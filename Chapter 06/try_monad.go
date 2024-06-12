package main

import (
 "fmt"
 "errors"
)

// Try is a monad that encapsulates a computation that might fail.
type Try struct {
 Result interface{}
 Err    error
}

// Bind applies a function to the result if there is no error.
func (t Try) Bind(f func(interface{}) Try) Try {
 if t.Err != nil {
  return Try{nil, t.Err}
 }
 return f(t.Result)
}

// ReturnTry lifts a result and error into the Try monad.
func ReturnTry(result interface{}, err error) Try {
 return Try{Result: result, Err: err}
}

// Example function that may fail.
func divide(a, b int) Try {
 if b == 0 {
  return ReturnTry(0, errors.New("division by zero"))
 }
 return ReturnTry(a / b, nil)
}

func main() {
 // Chaining operations with the Try monad.
 result := divide(10, 2).Bind(func(value interface{}) Try {
  return ReturnTry(value.(int)*2, nil)
 })

 if result.Err != nil {
  fmt.Println("Error:", result.Err)
 } else {
  fmt.Println("Result:", result.Result)
 }
}
