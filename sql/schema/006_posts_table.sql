-- +goose Up
CREATE TABLE posts (
	id UUID PRIMARY KEY,
	title TEXT NOT NULL,
	url TEXT NOT NULL UNIQUE,
	description TEXT,
	published_at TIMESTAMP,
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP NOT NULL,
	feed_id UUID NOT NULL REFERENCES feeds,
	FOREIGN KEY(feed_id) REFERENCES feeds(id)
);

-- +goose Down
DROP TABLE posts;
