package main

import (
	"net/http"
	"strconv"

	"github.com/Abi-Liu/rss-aggregator/internal/database"
)

func (c *apiConfig) getUserPosts(w http.ResponseWriter, r *http.Request, user database.User) {
	limit := r.URL.Query().Get("limit")
	num, err := strconv.Atoi(limit)
	if err != nil {
		num = 10
	}
	posts, err := c.DB.GetPostsByUser(r.Context(), database.GetPostsByUserParams{
		UserID: user.ID,
		Limit:  int32(num),
	})
	if err != nil {
		respondWithError(w, http.StatusNotFound, "No posts found, subscribe to a feed")
		return
	}

	respondWithJson(w, http.StatusOK, databasePostsToPosts(posts))
}
