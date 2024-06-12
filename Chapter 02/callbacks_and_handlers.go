type RouteHandler func(http.ResponseWriter, *http.Request)

func CreateRouteHandler(message string) RouteHandler {
    return func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprint(w, message)
    }
}

func main() {
    http.HandleFunc("/hello", CreateRouteHandler("Hello, world!"))
    http.HandleFunc("/greet", CreateRouteHandler("Greetings!"))
    
    http.ListenAndServe(":8080", nil)
}
