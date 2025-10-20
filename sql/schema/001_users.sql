-- +goose Up
CREATE TABLE users(
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP NOT NULL,
	name VARCHAR(255) NOT NULL
);

-- +goose Down
DROP TABLE USERS;
