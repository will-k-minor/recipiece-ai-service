package handlers

import (
    "encoding/json"
    "net/http"
    "recipiece-ai-service/clients"
)

func ChatHandler(w http.ResponseWriter, r *http.Request) {
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

    chatClient := clients.NewChatGPTClient()
    apiResponse, err := chatClient.CreateThread(message)
    if err != nil {
        http.Error(w, "Failed to create thread with ChatGPT API", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(apiResponse)
}
