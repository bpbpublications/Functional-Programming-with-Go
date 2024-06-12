func multiplyBy(factor int) func(int) int {
    return func(x int) int {
        return x * factor
    }
}

func main() {
    double := multiplyBy(2)
    result := double(5) // result will be 10
}
