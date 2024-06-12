package main

import (
    "fmt"
    "net/http"
    "time"
)

func fetchURL(url string) {
    start := time.Now()
    resp, err := http.Get(url)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Printf("Fetched %s in %v\n", url, time.Since(start))
    resp.Body.Close()
}

func main() {
    urls := []string{
        "https://www.google.com",
        "https://www.facebook.com",
        "https://www.twitter.com",
    }

    for _, url := range urls {
        go fetchURL(url)
    }

    time.Sleep(5 * time.Second) // wait for goroutines to finish
}
