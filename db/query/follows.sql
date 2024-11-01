-- name: CreateFollow :one
INSERT INTO follows (
  uid,
  following_uid
) VALUES (
  $1, $2
) RETURNING *;

-- name: DeleteFollow :one
DELETE FROM follows
WHERE uid = $1 AND following_uid = $2
RETURNING *;

-- name: GetFollow :one
SELECT * FROM follows
WHERE uid = $1 AND following_uid = $2;

-- name: GetFollowInfo :one
SELECT 
EXISTS (
        SELECT 1 
        FROM follows f2
        WHERE f2.uid = sqlc.arg(own_uid) AND f2.following_uid = $1
    ) AS is_following,
       COUNT(CASE WHEN f.uid = $1 THEN 1 ELSE NULL END) AS following_count,
       COUNT(CASE WHEN f.following_uid = $1 THEN 1 ELSE NULL END) AS followers_count
FROM follows f
WHERE f.uid = $1 or f.following_uid = $1;
