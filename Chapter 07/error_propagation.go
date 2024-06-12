package main

import (
    "errors"
    "fmt"
)

func performTask(id int) (string, error) {
    if id < 0 {
        return "", errors.New("invalid id")
    }
    return fmt.Sprintf("Result for %d", id), nil
}

func worker(id int, resultChan chan<- string, errChan chan<- error) {
    result, err := performTask(id)
    if err != nil {
        errChan <- err
        return
    }
    resultChan <- result
}

func main() {
    ids := []int{1, -2, 3}
    resultChan := make(chan string, len(ids))
    errChan := make(chan error, len(ids))

    for _, id := range ids {
        go worker(id, resultChan, errChan)
    }

    for range ids {
        select {
        case result := <-resultChan:
            fmt.Println(result)
        case err := <-errChan:
            fmt.Println("Error:", err)
        }
    }
}
