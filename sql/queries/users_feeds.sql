-- name: FollowFeed :one
INSERT INTO users_feeds (id, user_id, feed_id, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: DeleteFollowFeed :execrows
DELETE FROM users_feeds WHERE id = $1;

-- name: GetFeedById :one
SELECT * FROM users_feeds WHERE id = $1;

-- name: GetUsersFeeds :many
SELECT * FROM users_feeds WHERE user_id = $1;
