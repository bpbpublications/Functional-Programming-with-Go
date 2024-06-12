func Concurrently(tasks ...func() int) []int {
	var results []int
	ch := make(chan int)
   
	for _, task := range tasks {
	 go func(t func() int) {
	  ch <- t()
	 }(task)
	}
   
	for range tasks {
	 results = append(results, <-ch)
	}
	return results
   }
   
   func main() {
	tasks := []func() int{
	 func() int { return 1 + 1 },
	 func() int { return 2 + 2 },
	 func() int { return 3 + 3 },
	}
   
	fmt.Println(Concurrently(tasks...))
   }
   