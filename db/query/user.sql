-- name: CreateUser :one
INSERT INTO "user" (
  "uid",
  "username",
  "currency"
) VALUES (
  $1, $2, $3
) RETURNING *;  
  
-- name: GetUser :one
SELECT * FROM "user"
WHERE "uid" = $1;

-- name: GetUserByUsername :one
SELECT * FROM "user"
WHERE "username" = $1;

-- name: UpdateUser :one
UPDATE "user" SET
  "username" = COALESCE(sqlc.narg(username), "username"),
  "currency" = COALESCE(sqlc.narg(currency), "currency")
WHERE "uid" = sqlc.arg(uid)
RETURNING *;

-- name: GetFollowUsers :many
SELECT
"user"."user_id", 
"user"."username"
FROM "user"
JOIN "follow_user"
ON "user"."user_id" = "follow_user"."follow_user_id"
WHERE "follow_user"."user_id" = $1;  
