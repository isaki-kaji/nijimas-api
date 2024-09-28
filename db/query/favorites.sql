-- name: CreateFavorite :one
INSERT INTO favorites (
  post_id,
  uid
) VALUES (
  $1, $2
) RETURNING *;

-- name: DeleteFavorite :one
DELETE FROM favorites
WHERE post_id = $1 AND uid = $2
RETURNING *;

-- name: GetFavorite :one
SELECT * FROM favorites
WHERE post_id = $1 AND uid = $2;
