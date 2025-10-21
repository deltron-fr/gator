-- name: Createfeed :one

INSERT INTO feeds(id, name, url, created_at, updated_at, user_id)
VALUES(
	$1,
	$2,
	$3,
	$4,
	$5,
	$6
)
RETURNING *;

-- name: GetFeeds :many

SELECT * FROM feeds;

-- name: GetUserFeeds :one
SELECT users.name
FROM users
INNER JOIN feeds
ON users.id = $1;

-- name: CreateFeedFollow :one

WITH inserted_feed_follow AS (
	INSERT INTO feed_follows(id, created_at, updated_at, user_id, feed_id)
	VALUES(
		$1,
		$2,
		$3,
		$4,
		$5
	)
	RETURNING *
)

SELECT 
	inserted_feed_follow.*,
	feeds.name AS feed_name,
	users.name AS user_name
FROM inserted_feed_follow
INNER JOIN feeds ON feeds.id = inserted_feed_follow.feed_id
INNER JOIN users ON users.id = inserted_feed_follow.user_id;

-- name: GetFeed :one

SELECT *
FROM feeds
WHERE $1 = url;

-- name: GetFeedFollows :many

SELECT users.name AS user_name, feeds.name AS feed_name
FROM feed_follows
INNER JOIN feeds ON feeds.id = feed_follows.feed_id
INNER JOIN users ON users.id = feed_follows.user_id
WHERE feed_follows.user_id = $1;

-- name: DeleteFeedFollow :exec

DELETE FROM feed_follows 
WHERE ($1 = user_id) AND ($2 = feed_id);


-- name: MarkFeedFetched :exec

UPDATE feeds
SET last_fetched_at = NOW(), updated_at = NOW()
WHERE id = $1;

-- name: GetNextFeedToFetch :one

SELECT *
FROM feeds
ORDER BY last_fetched_at ASC NULLS FIRST LIMIT 1;

