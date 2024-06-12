package main

import (
 "fmt"
)

// Create a new bank account module with hidden details
func NewBankAccount(accountName string, initialBalance float64) func() (string, float64) {
 name := accountName
 balance := initialBalance

 // The closure provides access to account details
 return func() (string, float64) {
  return name, balance
 }
}

func main() {
 // Create two independent bank account instances with hidden details
 accountA := NewBankAccount("Savings Account", 1000.0)
 accountB := NewBankAccount("Checking Account", 500.0)

 // Access and display account details
 nameA, balanceA := accountA()
 nameB, balanceB := accountB()

 fmt.Printf("%s: Balance %.2f\n", nameA, balanceA) // Output: Savings Account: Balance 1000.00
 fmt.Printf("%s: Balance %.2f\n", nameB, balanceB) // Output: Checking Account: Balance 500.00
}
