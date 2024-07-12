-- name: CreatePost :one
INSERT INTO posts (id, title, url, description, published_at, created_at, updated_at, feed_id)
VALUES ($1, $2, $3, $4, $5, NOW(), NOW(), $6)
RETURNING *;

-- name: GetPostsByUser :many
SELECT posts.* FROM posts
INNER JOIN users_feeds
ON users_feeds.feed_id = posts.feed_id
WHERE users_feeds.user_id = $1
ORDER BY posts.published_at ASC
LIMIT $2;

