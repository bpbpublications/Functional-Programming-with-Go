package main

import (
 "fmt"
 "strings"
)

// Maybe represents a computation that might or might not return a value.
type Maybe struct {
 Value interface{}
}

// Bind applies a function to the value in the Maybe monad, if it exists.
func (m Maybe) Bind(f func(interface{}) Maybe) Maybe {
 if m.Value == nil {
  return m // Propagate the absence of a value
 }
 return f(m.Value)
}

// Return lifts a value into the Maybe monad.
func Return(value interface{}) Maybe {
 return Maybe{Value: value}
}

// Example function that may return a value or nil
func fetchGreeting(language string) Maybe {
 greetings := map[string]string{
  "English": "Hello",
  "Spanish": "Hola",
  "French":  "Bonjour",
 }
 if greeting, ok := greetings[language]; ok {
  return Return(greeting)
 }
 return Maybe{nil}
}

// Another example function to work with Maybe monad
func exclaim(input interface{}) Maybe {
 if str, ok := input.(string); ok {
  return Return(strings.ToUpper(str) + "!")
 }
 return Maybe{nil}
}

func main() {
 // Example usage of the Maybe monad
 result := fetchGreeting("English").Bind(exclaim)

 if result.Value == nil {
  fmt.Println("No value")
 } else {
  fmt.Printf("Success: %v\n", result.Value)
 }
}
