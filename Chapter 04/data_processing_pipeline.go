func FilterGreaterThan(minValue int) func([]int) []int {
    return func(data []int) []int {
        var result []int
        for _, v := range data {
            if v > minValue {
                result = append(result, v)
            }
        }
        return result
    }
}
