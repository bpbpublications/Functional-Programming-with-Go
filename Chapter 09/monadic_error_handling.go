package main

import (
    "encoding/json"
    "errors"
    "log"
    "net/http"
)

// Either represents a container that can hold either a value (Right) or an error (Left).
type Either struct {
    Value interface{}
    Err   error
}

// UserProfile represents a user's profile information.
type UserProfile struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
}

// fetchUserProfile simulates fetching a user profile from a database.
func fetchUserProfile(userID int) Either {
    // Simulate a database operation
    if userID == 1 {
        return Either{Value: UserProfile{ID: 1, Name: "Rex Solaris"}, Err: nil}
    }
    return Either{Value: nil, Err: errors.New("User not found")} // Simulate an error
}

// userProfileHandler is an HTTP handler that uses the Either structure to process the request.
func userProfileHandler(w http.ResponseWriter, r *http.Request) {
    // Simulate fetching a user profile for user with ID 1
    result := fetchUserProfile(1)

    if result.Err != nil {
        http.Error(w, result.Err.Error(), http.StatusInternalServerError)
        return
    }

    userProfile := result.Value.(UserProfile)
    json.NewEncoder(w).Encode(userProfile)
}

func main() {
    http.HandleFunc("/user-profile", userProfileHandler)
    log.Println("Server started on :8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
}
