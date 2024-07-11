package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Abi-Liu/rss-aggregator/internal/database"
	"github.com/google/uuid"
)

func (c *apiConfig) createFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}

	decoder := json.NewDecoder(r.Body)

	params := &parameters{}
	err := decoder.Decode(params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to decode parameters")
		return
	}

	feed, err := c.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		Url:       params.URL,
		UserID:    user.ID,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	})

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to create feed: "+err.Error())
		return
	}

	userFeed, err := c.DB.FollowFeed(r.Context(), database.FollowFeedParams{
		ID:        uuid.New(),
		UserID:    user.ID,
		FeedID:    feed.ID,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	})

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to follow feed")
		return
	}

	type response struct {
		Feed       Feed     `json:"feed"`
		FeedFollow UserFeed `json:"feed_follow"`
	}

	respondWithJson(w, http.StatusCreated, response{
		Feed:       databaseFeedToFeed(feed),
		FeedFollow: databaseUserFeedToUserFeed(userFeed),
	})
}

func (c *apiConfig) getAllFeeds(w http.ResponseWriter, r *http.Request) {
	feeds, err := c.DB.GetAllFeeds(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Could not retrieve feeds")
		return
	}

	respondWithJson(w, http.StatusOK, databaseFeedsToFeeds(feeds))
}
