package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Abi-Liu/rss-aggregator/internal/database"
	"github.com/google/uuid"
)

func (c *apiConfig) followFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedId uuid.UUID `json:"feed_id"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to decode parameters")
		return
	}

	userFeed, err := c.DB.FollowFeed(r.Context(), database.FollowFeedParams{
		ID:        uuid.New(),
		UserID:    user.ID,
		FeedID:    params.FeedId,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	})

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to follow feed")
		return
	}

	respondWithJson(w, http.StatusOK, databaseUserFeedToUserFeed(userFeed))
}

func (c *apiConfig) deleteFeedFollow(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("feedFollowID")
	uuid, err := uuid.Parse(id)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Could not parse uuid: "+err.Error())
		return
	}

	_, err = c.DB.DeleteFollowFeed(r.Context(), uuid)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Could not delete user follow")
		return
	}

	respondWithJson(w, http.StatusNoContent, "")
}

func (c *apiConfig) getUsersFeeds(w http.ResponseWriter, r *http.Request, user database.User) {
	feeds, err := c.DB.GetUsersFeeds(r.Context(), user.ID)

	if err != nil {
		respondWithError(w, http.StatusNotFound, "No followed feeds")
		return
	}

	respondWithJson(w, http.StatusOK, databaseUserFeedsToUserFeeds(feeds))
}
