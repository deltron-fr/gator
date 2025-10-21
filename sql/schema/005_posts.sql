-- +goose Up

CREATE TABLE posts(
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	created_at  TIMESTAMP NOT NULL,
	updated_at TIMESTAMP NOT NULL,
	title VARCHAR(255) NOT NULL,
	url VARCHAR(255) UNIQUE NOT NULL,
	description VARCHAR(255) NOT NULL,
	published_at VARCHAR(255),
	feed_id UUID NOT NULL,
	CONSTRAINT fk_posts
		FOREIGN KEY (feed_id)
		REFERENCES feeds (id)
);

-- +goose Down

DROP TABLE posts;

