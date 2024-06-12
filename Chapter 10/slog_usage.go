package main

import (
 "log/slog"
 "math"
 "os"
)

// IsPrime checks if a number is prime, logging the process.
func IsPrime(n int) bool {
 if n < 2 {
  slog.Info("Checked number is less than 2, not prime", "number", n)
  return false
 }
 for i := 2; float64(i) <= math.Sqrt(float64(n)); i++ {
  if n%i == 0 {
   slog.Info("Found a divisor, not prime", "number", n, "divisor", i)
   return false
  }
 }
 slog.Info("Number is prime", "number", n)
 return true
}

func main() {
 logger := slog.New(slog.NewJSONHandler(os.Stderr,
  &slog.HandlerOptions{Level: slog.LevelDebug}))
 slog.SetDefault(logger)
 _ = IsPrime(29)
 _ = IsPrime(50)
}
