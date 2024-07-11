package main

import (
	"time"

	"github.com/Abi-Liu/rss-aggregator/internal/database"
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	ApiKey    string    `json:"api_key"`
}

func databaseUserToUser(dbUser database.User) User {
	return User{
		ID:        dbUser.ID,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		Name:      dbUser.Name,
		ApiKey:    dbUser.ApiKey,
	}
}

type Feed struct {
	ID          uuid.UUID  `json:"id"`
	URL         string     `json:"url"`
	UserId      uuid.UUID  `json:"user_id"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	LastFetched *time.Time `json:"last_fetched"`
}

func databaseFeedToFeed(dbFeed database.Feed) Feed {
	return Feed{
		ID:          dbFeed.ID,
		URL:         dbFeed.Url,
		UserId:      dbFeed.UserID,
		CreatedAt:   dbFeed.CreatedAt,
		UpdatedAt:   dbFeed.UpdatedAt,
		LastFetched: &dbFeed.LastFetched.Time,
	}
}

func databaseFeedsToFeeds(feeds []database.Feed) []Feed {
	res := make([]Feed, len(feeds))
	for i, v := range feeds {
		res[i] = databaseFeedToFeed(v)
	}
	return res
}

type UserFeed struct {
	ID        uuid.UUID `json:"id"`
	UserId    uuid.UUID `json:"user_id"`
	FeedId    uuid.UUID `json:"feed_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func databaseUserFeedToUserFeed(uf database.UsersFeed) UserFeed {
	return UserFeed{
		ID:        uf.ID,
		UserId:    uf.UserID,
		FeedId:    uf.FeedID,
		CreatedAt: uf.CreatedAt,
		UpdatedAt: uf.UpdatedAt,
	}
}

func databaseUserFeedsToUserFeeds(userFeeds []database.UsersFeed) []UserFeed {
	res := make([]UserFeed, len(userFeeds))
	for i, v := range userFeeds {
		res[i] = databaseUserFeedToUserFeed(v)
	}
	return res
}

type Post struct {
	ID          uuid.UUID  `json:"id"`
	Title       string     `json:"title"`
	Url         string     `json:"url"`
	Description *string    `json:"description"`
	PublishedAt *time.Time `json:"published_at"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	FeedId      uuid.UUID  `json:"feed_id"`
}

func databasePostToPost(post database.Post) Post {
	return Post{
		ID:          post.ID,
		Title:       post.Title,
		Url:         post.Url,
		Description: &post.Description.String,
		PublishedAt: &post.PublishedAt.Time,
		CreatedAt:   post.CreatedAt,
		UpdatedAt:   post.UpdatedAt,
		FeedId:      post.FeedID,
	}
}

func databasePostsToPosts(posts []database.Post) []Post {
	res := make([]Post, len(posts))
	for i, v := range posts {
		res[i] = databasePostToPost(v)
	}
	return res
}
