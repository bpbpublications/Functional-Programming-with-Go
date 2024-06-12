package main

import (
 "fmt"
 "errors"
)

// Either represents a computation that results in either a correct value (Right) or an error (Left).
type Either struct {
 Right interface{}
 Left  error
}

// Bind applies a function to the Right value if there is no error.
func (e Either) Bind(f func(interface{}) Either) Either {
 if e.Left != nil {
  return e // Propagate the error
 }
 return f(e.Right)
}

// ReturnEither lifts a result into the Either monad as a Right value.
func ReturnEither(result interface{}) Either {
 return Either{Right: result, Left: nil}
}

// ReturnEitherError lifts an error into the Either monad as a Left value.
func ReturnEitherError(err error) Either {
 return Either{Right: nil, Left: err}
}

// Example operation that returns Either.
func fetchData(key string) Either {
 data := map[string]string{"key1": "value1", "key2": "value2"}
 if value, ok := data[key]; ok {
  return ReturnEither(value)
 }
 return ReturnEitherError(errors.New("key not found"))
}

func main() {
 // Chaining operations with the Either monad.
 result := fetchData("key1").Bind(func(value interface{}) Either {
  // Further processing with the fetched data
  return ReturnEither(value.(string) + " processed")
 })

 if result.Left != nil {
  fmt.Println("Error:", result.Left)
 } else {
  fmt.Println("Success:", result.Right)
 }
}
