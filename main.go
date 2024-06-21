package main

import (
    "encoding/json"
    "log"
    "net/http"
    "github.com/go-chi/chi/v5"
    "github.com/go-chi/chi/v5/middleware"
)

func main() {
    router := chi.NewRouter()
    router.Use(middleware.Logger)

    router.Post("/hello", helloHandler)

    log.Fatal(http.ListenAndServe(":8080", router))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
    var requestBody map[string]string
    err := json.NewDecoder(r.Body).Decode(&requestBody)
    if err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    message, exists := requestBody["message"]
    if !exists {
        http.Error(w, "Message field is required", http.StatusBadRequest)
        return
    }

    if message == "Hello" {
        json.NewEncoder(w).Encode(map[string]string{"response": "World"})
    } else {
        json.NewEncoder(w).Encode(map[string]string{"response": "Unknown message"})
    }
}

