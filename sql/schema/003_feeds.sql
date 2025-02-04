-- +goose Up
CREATE TABLE feeds (
	id UUID PRIMARY KEY,
	url TEXT UNIQUE NOT NULL,
	user_id UUID NOT NULL REFERENCES users ON DELETE CASCADE,
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP NOT NULL,
	FOREIGN KEY (user_id) REFERENCES users(id)
);

-- +goose Down
DROP TABLE feeds;
