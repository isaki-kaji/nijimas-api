-- name: GetUser :one
SELECT * FROM "user"
WHERE "uid" = $1;