package main

import (
 "fmt"
 "net/http"
 "sync"
 "time"

 _ "net/http/pprof"
)

func toughGrind(waitGroup *sync.WaitGroup) {
 defer waitGroup.Done()
 fmt.Printf("Goroutine starts at %v\n", time.Now())

 a := []int{}
 for i := 0; i < 500000; i++ {
  a = append(a, i)
  if (i % 10000) == 0 {
   time.Sleep(500 * time.Millisecond)
  }
 }

 time.Sleep(2 * time.Second)
 fmt.Printf("Goroutine ends at %v\n", time.Now())
}

func main() {
 var waitGroup sync.WaitGroup

 go func() {
  // Server for pprof
  http.ListenAndServe("localhost:8080", nil)
 }()
 waitGroup.Add(1) // This does not let program exit
 waitGroup.Add(1) // for the Goroutine toughGrind
 go toughGrind(&waitGroup)
 waitGroup.Wait()
}
