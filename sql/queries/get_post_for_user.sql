-- name: GetPostsForUser :many
SELECT
    posts.id,
    posts.created_at,
    posts.updated_at,
    posts.title,
    posts.url,
    posts.description,
    posts.published_at,
    posts.feed_id
FROM
    posts
    INNER JOIN feeds ON posts.feed_id = feeds.id
    INNER JOIN feed_follows ON feeds.id = feed_follows.feed_id
    INNER JOIN users ON feed_follows.user_id = users.id
WHERE
    users.id = $1
ORDER BY
    posts.published_at DESC
LIMIT $2;
