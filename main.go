package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"recipiece-ai-service/handlers"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	router := chi.NewRouter()
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
	})
	router.Use(c.Handler)
	router.Use(middleware.Logger)

	router.Post("/hello", handlers.HelloHandler)
	router.Post("/threads/create", handlers.CreateThread)
	router.Post("/threads/{threadId}/send-message", handlers.SendMessageToThread)
	router.Post("/threads/{threadId}/run", handlers.RunAssistant)
	router.Post("/threads/{threadId}/runs/{runId}/cancel", handlers.CancelRun)
	router.Get("/threads", handlers.ListThreads)
	router.Get("/threads/{threadId}/messages", handlers.GetMessagesFromThread)
	router.Get("/threads/{threadId}/runs/{runId}", handlers.GetRunDetails)
	router.Post("/threads/{threadId}/runs/{runId}/submit-tools-output", handlers.SubmitTools)

	log.Fatal(http.ListenAndServe(":8080", router))
}
