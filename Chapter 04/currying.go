package main

import (
 "fmt"
 "strings"
)

// Function to append prefix
func AppendPrefix(prefix string) func(string) string {
 return func(baseStr string) string {
  return prefix + baseStr
 }
}

// Function to replace characters
func ReplaceChars(oldChar, newChar rune) func(string) string {
 return func(baseStr string) string {
  return strings.ReplaceAll(baseStr, string(oldChar), string(newChar))
 }
}

// Function to change case
func ToUpperCase() func(string) string {
 return func(baseStr string) string {
  return strings.ToUpper(baseStr)
 }
}

func main() {
 // Create a function to append "GoLang: " prefix
 addPrefix := AppendPrefix("GoLang: ")

 // Create a function to replace 'a' with '@'
 replaceChars := ReplaceChars('a', '@')

 // Create a function to convert to uppercase
 toUpper := ToUpperCase()

 // Original string
 originalStr := "functional programming"

 // Applying the curried functions
 result := toUpper(replaceChars(addPrefix(originalStr)))

 fmt.Println(result) // Output: "GOLANG: FUNCTION@L PROGR@MMING"
}
