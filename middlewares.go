package main

import (
	"net/http"

	"github.com/Abi-Liu/rss-aggregator/internal/auth"
	"github.com/Abi-Liu/rss-aggregator/internal/database"
)

type AuthHandler func(http.ResponseWriter, *http.Request, database.User)

func (c *apiConfig) middlewareAuth(handler AuthHandler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		key, err := auth.GetApiKey(r.Header)
		if err != nil {
			respondWithError(w, http.StatusUnauthorized, "Unauthorized, Please provide an APIKEY")
			return
		}
		user, err := c.DB.GetUserByApiKey(r.Context(), key)
		if err != nil {
			respondWithError(w, http.StatusNotFound, "User not found")
			return
		}

		handler(w, r, user)
	})
}
