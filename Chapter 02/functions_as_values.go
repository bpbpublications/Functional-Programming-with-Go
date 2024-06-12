func add(a, b int) int {
    return a + b
}

func main() {
    // Assigning a function to a variable
    operation := add

    // Using the function variable
    result := operation(3, 4) // result will be 7
}
