// Function to sum numbers in a slice and send the result on a channel
func sum(numbers []int, resultChan chan int) {
	sum := 0
	for _, number := range numbers {
	 sum += number
	}
	resultChan <- sum // send the sum to the channel
   }
   
   func main() {
	numbers := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20} // a slice of integers
	resultChan := make(chan int)                        // create a channel to communicate results
   
	// Split the slice into two halves
	mid := len(numbers) / 2
   
	// Start two goroutines to sum each half of the array
	go sum(numbers[:mid], resultChan) // first half
	go sum(numbers[mid:], resultChan) // second half
   
	// Receive results from both goroutines
	sum1 := <-resultChan // receive from first goroutine
	sum2 := <-resultChan // receive from second goroutine
   
	// Calculate the total sum of all numbers
	totalSum := sum1 + sum2
	fmt.Println("Total Sum:", totalSum)
   }
   