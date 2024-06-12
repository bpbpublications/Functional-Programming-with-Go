func ApplyFunction(f func(int) int, value int) int {
    return f(value)
}

increment := func(x int) int { return x + 1 }
result := ApplyFunction(increment, 4)  // 5
