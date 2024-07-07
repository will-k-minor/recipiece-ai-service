package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
	"os"
	"recipiece-ai-service/clients"
	"recipiece-ai-service/utils"
)

func ListThreads(w http.ResponseWriter, r *http.Request) {
	threads, err := utils.ReadFileLines("threads.txt")
	if err != nil {
		http.Error(w, "Something went wrong getting threads", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(threads)
}

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

	// Write results to a log file
	file, err := os.OpenFile("threads.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		http.Error(w, "Failed to open log file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	logEntry := fmt.Sprintf("%v\n", apiResponse["id"])
	if _, err := file.WriteString(logEntry); err != nil {
		http.Error(w, "Failed to write to log file", http.StatusInternalServerError)
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

	threadId := chi.URLParam(r, "threadId")
	if len(threadId) == 0 {
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

	threadId := chi.URLParam(r, "threadId")
	if len(threadId) == 0 {
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

func CancelRun(w http.ResponseWriter, r *http.Request) {
	threadId := chi.URLParam(r, "threadId")
	runId := chi.URLParam(r, "runId")

	chatClient := clients.NewChatGPTClient()
	apiResponse, err := chatClient.CancelRun(threadId, runId)
	if err != nil {
		http.Error(w, "Failed to cancel run with ChatGPT API", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(apiResponse)
}

func GetRunDetails(w http.ResponseWriter, r *http.Request) {
	threadId := chi.URLParam(r, "threadId")
	runId := chi.URLParam(r, "runId")

	chatClient := clients.NewChatGPTClient()
	apiResponse, err := chatClient.GetRunDetails(threadId, runId)
	if err != nil {
		http.Error(w, "Failed to get run details with ChatGPT API", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(apiResponse)
}

// Some crappy way to unjam runs when "required_action" appears
func SubmitTools(w http.ResponseWriter, r *http.Request) {
	threadId := chi.URLParam(r, "threadId")
	runId := chi.URLParam(r, "runId")
	var requestBody map[string]string
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	toolId, exists := requestBody["toolId"]
	if !exists {
		http.Error(w, "Tools field is required", http.StatusBadRequest)
		return
	}

	output, exists := requestBody["output"]
	if !exists {
		http.Error(w, "Output field is required", http.StatusBadRequest)
		return
	}

	chatClient := clients.NewChatGPTClient()
	apiResponse, err := chatClient.SubmitToolOutput(threadId, runId, toolId, output)
	if err != nil {
		http.Error(w, "Failed to submit tools with ChatGPT API", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(apiResponse)
}
