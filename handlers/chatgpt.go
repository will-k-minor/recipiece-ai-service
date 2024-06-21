package handlers

import (
	"encoding/json"
	"net/http"
	"recipiece-ai-service/clients"
    "github.com/go-chi/chi/v5"
)

func CreateThread(w http.ResponseWriter, r *http.Request) {
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

func SendMessageToThread(w http.ResponseWriter, r *http.Request) {
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

	threadId, exists := requestBody["threadId"]
	if !exists {
		http.Error(w, "Thread ID field is required", http.StatusBadRequest)
		return
	}

	chatClient := clients.NewChatGPTClient()
	apiResponse, err := chatClient.CreateMessage(threadId, message)
	if err != nil {
		http.Error(w, "Failed to send message to thread with ChatGPT API", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(apiResponse)
}

func RunAssistant(w http.ResponseWriter, r *http.Request) {
	var requestBody map[string]string
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	threadId, exists := requestBody["threadId"]
	if !exists {
		http.Error(w, "Message field is required", http.StatusBadRequest)
		return
	}

	chatClient := clients.NewChatGPTClient()
	apiResponse, err := chatClient.RunAssistant(threadId)
	if err != nil {
		http.Error(w, "Failed to run assistant with ChatGPT API", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(apiResponse)
}

func GetMessagesFromThread(w http.ResponseWriter, r *http.Request) {
	threadId := chi.URLParam(r, "threadId")

	chatClient := clients.NewChatGPTClient()
	apiResponse, err := chatClient.GetMessagesFromThread(threadId)
	if err != nil {
		http.Error(w, "Failed to get messages from thread with ChatGPT API", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(apiResponse)
}