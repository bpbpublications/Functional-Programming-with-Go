func Counter() func() int {
    count := 0
    return func() int {
        count++
        return count
    }
}

func main() {
    counterA := Counter()
    counterB := Counter()

    fmt.Println("Counter A:", counterA()) // Output: Counter A: 1
    fmt.Println("Counter A:", counterA()) // Output: Counter A: 2
    fmt.Println("Counter B:", counterB()) // Output: Counter B: 1
    fmt.Println("Counter A:", counterA()) // Output: Counter A: 3
}
