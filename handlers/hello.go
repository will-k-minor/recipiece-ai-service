package handlers

import (
	"encoding/json"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
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