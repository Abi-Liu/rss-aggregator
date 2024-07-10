package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Abi-Liu/rss-aggregator/internal/database"
	"github.com/google/uuid"
)

func (c *apiConfig) createUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(r.Body)

	params := &parameters{}
	err := decoder.Decode(params)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Failed to decode parameters: %s", err))
		return
	}

	createUserArgs := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	}
	user, err := c.DB.CreateUser(r.Context(), createUserArgs)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Failed to create user: %s", err))
		return
	}

	respondWithJson(w, http.StatusCreated, databaseUserToUser(user))
}
