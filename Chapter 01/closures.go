func greetWithPrefix(prefix string) func(string) { 
	return func(name string) {
	  fmt.Println(prefix, name)
	} 
  }
  helloGreeting := greetWithPrefix("Hello")
  helloGreeting("Alice") // Outputs: Hello Alice 
  