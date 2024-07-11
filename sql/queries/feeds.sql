-- name: CreateFeed :one
INSERT INTO feeds (id, url, user_id) VALUES ($1, $2, $3) RETURNING *;
