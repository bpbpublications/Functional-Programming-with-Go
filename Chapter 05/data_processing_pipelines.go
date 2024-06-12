func processData(data []int) []int {
    result := make([]int, len(data))
    for i, v := range data {
        result[i] = v * v // Squares each element
    }
    return result
}

func main() {
    originalData := []int{1, 2, 3, 4}
    processedData := processData(originalData)
    fmt.Println(processedData) // Output: [1 4 9 16]
}
