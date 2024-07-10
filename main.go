package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Failed to load environment variables: %v\n", err)
	}
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("Port env variable is not set")
	}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /v1/healthz", getHealthStatus)
	mux.HandleFunc("GET /v1/err", simulateError)

	server := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Printf("Starting server on port %s", port)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Failed to start on port %s: %v\n", port, err)
	}
}
