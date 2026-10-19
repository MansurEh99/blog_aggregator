-- name: GetFeeds :many
SELECT
    feeds.name,
    feeds.url,
    users.name AS user_name
FROM
    feeds
    INNER JOIN users ON feeds.User_id = users.id;
