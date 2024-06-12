// Function to modify a number
func doubleValue(number int) {
	number *= 2
	fmt.Println("Inside doubleValue:", number)
   }
   
   func main() {
	original := 10
	fmt.Println("Before doubleValue:", original)
	doubleValue(original)
	fmt.Println("After doubleValue:", original)
   }
   