package main

import (
    "database/sql"
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    _ "github.com/go-sql-driver/mysql"
)

// ToDoItem represents a To-Do item.
type ToDoItem struct {
    ID   int    `json:"id"`
    Task string `json:"task"`
}

var db *sql.DB

func initDB() {
    var err error
    // Replace "user:password@/todoDB" with actual database connection information
    db, err = sql.Open("mysql", "user:password@/todoDB")
    if err != nil {
        log.Fatal(err)
    }

    if err = db.Ping(); err != nil {
        log.Fatal(err)
    }
}

// addToDo adds a new To-Do item to the database.
func addToDo(task string) (int64, error) {
    result, err := db.Exec("INSERT INTO todo_items (task) VALUES (?)", task)
    if err != nil {
        return 0, err
    }
    return result.LastInsertId()
}

// fetchToDos retrieves all To-Do items from the database.
func fetchToDos() ([]ToDoItem, error) {
    rows, err := db.Query("SELECT id, task FROM todo_items")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var todos []ToDoItem
    for rows.Next() {
        var todo ToDoItem
        if err := rows.Scan(&todo.ID, &todo.Task); err != nil {
            return nil, err
        }
        todos = append(todos, todo)
    }
    return todos, nil
}

func addHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != "POST" {
        http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
        return
    }

    var todo ToDoItem
    if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    id, err := addToDo(todo.Task)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    fmt.Fprintf(w, "Added To-Do item with ID: %d\n", id)
}

func fetchHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != "GET" {
        http.Error(w, "Only GET method is allowed", http.StatusMethodNotAllowed)
        return
    }

    todos, err := fetchToDos()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(todos)
}

func main() {
    initDB()

    http.HandleFunc("/add", addHandler)
    http.HandleFunc("/fetch", fetchHandler)

    fmt.Println("Server started on port 8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal("Failed to start server:", err)
    }
}
