func computeNewState(state int, input int) int {
    return state + input
}

func main() {
    currentState := 0
    newState := computeNewState(currentState, 5)
    fmt.Println(newState) // Output: 5
    // currentState remains unchanged
}
