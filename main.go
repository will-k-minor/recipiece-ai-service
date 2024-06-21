package main

import (
    "log"
    "net/http"
    "github.com/go-chi/chi/v5"
    "github.com/go-chi/chi/v5/middleware"
    "recipiece-ai-service/handlers"
    "github.com/joho/godotenv"
)

func main() {
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file");
    }

    router := chi.NewRouter();
    router.Use(middleware.Logger);

    router.Post("/hello", handlers.HelloHandler);
    router.Post("/threads/create", handlers.CreateThread);
    router.Post("/threads/{threadId}/send-message", handlers.SendMessageToThread);
    router.Post("/threads/{threadId}/run", handlers.RunAssistant);
    router.Post("/threads/{threadId}/runs/{runId}/cancel", handlers.CancelRun);
    router.Get("/threads/{threadId}/messages", handlers.GetMessagesFromThread);
    router.Get("/threads/{threadId}/runs/{runId}", handlers.GetRunDetails);
    router.Post("/threads/{threadId}/runs/{runId}/submit-tools-output", handlers.SubmitTools);

    log.Fatal(http.ListenAndServe(":8080", router));
}


