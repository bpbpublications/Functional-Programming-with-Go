package main

import (
    "encoding/json"
    "net/http"
)

// User represents an immutable user struct.
type User struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
}

// updateUser creates a new User instance with updated data. Does not alter the original user unchanged.
func updateUser(u User, newName string) User {
    // Create and return a new User instance with the updated name.
    return User{ID: u.ID, Name: newName}
}

func userHandler(w http.ResponseWriter, r *http.Request) {
    // Simulate fetching a user from a database or external source.
    originalUser := User{ID: 1, Name: "John Doe"}

    // Update the user's name while preserving immutability.
    updatedUser := updateUser(originalUser, "Jane Doe")

    // Marshal the updated user to JSON for response.
    jsonResponse, err := json.Marshal(updatedUser)
    if err != nil {
        http.Error(w, "Error creating JSON response", http.StatusInternalServerError)
        return
    }

    // Set the Content-Type header and write the JSON response.
    w.Header().Set("Content-Type", "application/json")
    w.Write(jsonResponse)
}

func main() {
    http.HandleFunc("/updateUser", userHandler)
    http.ListenAndServe(":8080", nil)
}
