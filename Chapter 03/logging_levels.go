func Logger(level string) func(string) {
    if level == "debug" {
        return func(msg string) { fmt.Println("[DEBUG]", msg) }
    }
    return func(msg string) { fmt.Println("[INFO]", msg) }
}

log := Logger("debug")
log("This is a debug message.")
