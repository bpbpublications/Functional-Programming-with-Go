func Reduce(slice []int, initializer int, op func(int, int) int) int {
    acc := initializer
    for _, v := range slice {
        acc = op(acc, v)
    }
    return acc
}

sum := Reduce([]int{1, 2, 3, 4}, 0, func(a, b int) int { return a + b })  // 10
