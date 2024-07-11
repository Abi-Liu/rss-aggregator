package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/Abi-Liu/rss-aggregator/internal/database"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	env, err := loadEnv()
	if err != nil {
		log.Fatalf("Error loading environment variables: %s\n", err)
	}

	db, err := sql.Open("postgres", env.dbString)
	if err != nil {
		log.Fatalf("Failed to open database connection: %s", err)
	}

	dbQueries := database.New(db)

	cfg := &apiConfig{
		DB: dbQueries,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /v1/healthz", getHealthStatus)
	mux.HandleFunc("GET /v1/err", simulateError)
	mux.HandleFunc("POST /v1/users", cfg.createUser)
	mux.HandleFunc("GET /v1/users", cfg.middlewareAuth(cfg.getCurrentUser))
	mux.HandleFunc("POST /v1/feeds", cfg.middlewareAuth(cfg.createFeed))
	mux.HandleFunc("GET /v1/feeds", cfg.getAllFeeds)
	mux.HandleFunc("POST /v1/feed_follows", cfg.middlewareAuth(cfg.followFeed))
	mux.HandleFunc("DELETE /v1/feed_follows/{feedFollowID}", cfg.deleteFeedFollow)

	server := &http.Server{
		Addr:    ":" + env.port,
		Handler: mux,
	}

	log.Printf("Starting server on port %s", env.port)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Failed to start on port %s: %v\n", env.port, err)
	}
}
