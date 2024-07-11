-- name: CreateFeed :one
INSERT INTO feeds (id, url, user_id, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;


-- name: GetAllFeeds :many
SELECT * FROM feeds;


-- name: GetNextFeedsToFetch :many
SELECT * FROM feeds
ORDER BY last_fetched DESC NULLS FIRST
LIMIT $1;

-- name: MarkFeedFetched :execrows
UPDATE feeds
SET last_fetched = NOW(), updated_at = NOW()
WHERE id = $1;
