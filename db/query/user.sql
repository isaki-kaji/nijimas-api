-- name: CreateUser :one
INSERT INTO "user" (
  "uid",
  "username",
  "country_code"
) VALUES (
  $1, $2, $3
) RETURNING *;  
  
-- name: GetUser :one
SELECT * FROM "user"
WHERE "uid" = $1;

-- name: UpdateUser :one
UPDATE "user" SET
  "username" = COALESCE(sqlc.narg(username), "username"),
  "profile_image_url" = COALESCE(sqlc.narg(profile_image_url), "profile_image_url"),
  "banner_image_url" = COALESCE(sqlc.narg(banner_image_url), "banner_image_url")
WHERE "uid" = sqlc.arg(uid)
RETURNING *;

-- name: GetFollowUsers :many
SELECT
"user"."uid", 
"user"."username"
FROM "user"
JOIN "follow_user"
ON "user"."uid" = "follow_user"."follow_user_id"
WHERE "follow_user"."uid" = $1;  
