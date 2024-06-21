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
    router.Post("/threads/create", handlers.ChatHandler);

    log.Fatal(http.ListenAndServe(":8080", router));
}


