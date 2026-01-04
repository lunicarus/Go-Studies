package main

import (
	"log"
	"net/http"
	"time"

	"GranmaCakesAPI/internal/server"
)

func main() {
	srv := server.New()

	httpServer := &http.Server{
		Addr:         ":8080",
		Handler:      srv.Router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	log.Println("Cakes Baking on :8080")
	log.Fatal(httpServer.ListenAndServe())
}
