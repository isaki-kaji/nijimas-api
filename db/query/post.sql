-- name: CreatePost :one
INSERT INTO "post" (
  "user_id",
  "main_category",
  "post_text",
  "photo_url",
  "room_id",
  "meal_flag",
  "location",
  "public_type_no"
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8
) RETURNING *;  
  
-- name: GetPostById :one
SELECT 
  "user"."username",
  "post"."main_category",
  "post".
FROM "post"
  


-- -- name: GetUserByUsername :one
-- SELECT * FROM "user"
-- WHERE "username" = $1;

-- -- name: UpdateUser :one
-- UPDATE "user" SET
--   "username" = COALESCE(sqlc.narg(username), "username"),
--   "currency" = COALESCE(sqlc.narg(currency), "currency")
-- WHERE "uid" = sqlc.arg(uid)
-- RETURNING *;

-- -- name: GetUsersByRoomID :many
-- SELECT
-- "user"."user_id", 
-- "user"."username"
-- FROM "user"
-- JOIN "follow_room"
-- ON "user"."user_id" = "follow_room"."user_id"
-- WHERE "follow_room"."room_id" = $1;  
