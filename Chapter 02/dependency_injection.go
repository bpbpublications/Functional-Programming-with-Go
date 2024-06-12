type Logger func(string)

func ProcessData(data string, log Logger) {
    // Process the data
    log("Data processed successfully")
}

func main() {
    // Define a logger function
    logToConsole := func(message string) {
        fmt.Println("Log:", message)
    }

    // Inject the logger into the data processing function
    ProcessData("Some data", logToConsole)
}
