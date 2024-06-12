func apply(operation func(int, int) int, a, b int) int {
    return operation(a, b)
}

func add(a, b int) int {
    return a + b
}

func main() {
    result := apply(add, 3, 4) // result will be 7
}
