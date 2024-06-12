package main

import (
 "fmt"
)

// min returns the smallest of three integers.
func min(a, b, c int) int {
 if a < b {
  if a < c {
   return a
  }
  return c
 }
 if b < c {
  return b
 }
 return c
}

// levenshteinDistance uses memoization to calculate the Levenshtein distance between two strings.
func levenshteinDistance(s1, s2 string, memo map[string]int) int {
 key := fmt.Sprintf("%d,%d", len(s1), len(s2))

 // Check if the result is in the memoization map.
 if val, found := memo[key]; found {
  return val
 }

 // Base cases for recursion.
 if len(s1) == 0 {
  return len(s2)
 }
 if len(s2) == 0 {
  return len(s1)
 }

 // If last characters of two strings are the same, ignore the last char
 // and recur for remaining string.
 if s1[len(s1)-1] == s2[len(s2)-1] {
  memo[key] = levenshteinDistance(s1[:len(s1)-1], s2[:len(s2)-1], memo)
 } else {
  // If the last character is different, consider all possibilities and
  // find the minimum.
  insert := levenshteinDistance(s1, s2[:len(s2)-1], memo)
  delete := levenshteinDistance(s1[:len(s1)-1], s2, memo)
  replace := levenshteinDistance(s1[:len(s1)-1], s2[:len(s2)-1], memo)
  memo[key] = 1 + min(insert, delete, replace)
 }

 return memo[key]
}

func main() {
 s1 := "kitten"
 s2 := "settings"
 memo := make(map[string]int) // Create a map to hold memoized results.

 distance := levenshteinDistance(s1, s2, memo)
 fmt.Printf("The Levenshtein distance between '%s' and '%s' is %d.\n", s1, s2, distance)
}
