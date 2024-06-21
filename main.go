package main

import (
    "log"
    "net/http"
    "github.com/go-chi/chi/v5"
    "github.com/go-chi/chi/v5/middleware"
    "recipiece-ai-service/handlers"
)

func main() {
    router := chi.NewRouter()
    router.Use(middleware.Logger)

    router.Post("/hello", handlers.HelloHandler)

    log.Fatal(http.ListenAndServe(":8080", router))
}


