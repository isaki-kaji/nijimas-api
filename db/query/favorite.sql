-- name: CreateFavorite :one
INSERT INTO "favorite" (
 "post_id",
 "uid"
) VALUES (
 $1,$2
) RETURNING *;

-- name: DeleteFavorite :one
DELETE FROM "favorite"
WHERE "post_id" = $1 AND "uid" = $2
RETURNING *;

-- name: GetFavorite :one
SELECT * FROM "favorite"
WHERE "post_id" = $1 AND "uid" = $2;