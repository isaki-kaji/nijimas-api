-- name: CreatePost :one
INSERT INTO "post" (
  "user_id",
  "main_category",
  "post_text",
  "photo_url",
  "meal_flag",
  "location",
  "public_type_no"
) VALUES (
  $1, $2, $3, $4, $5, $6, $7
) RETURNING *;  
  
-- name: GetPostById :one
SELECT 
  u."username",
  p."main_category",
  ps1."sub_category",
  ps2."sub_category",
  p."post_text",
  p."photo_url",
  p."location",
  p."public_type_no"
FROM "post" AS p
JOIN "user" AS u ON p."user_id" = u."user_id"
LEFT JOIN "post_subcategory" AS ps1
ON p."post_id" = ps1."post_id" AND ps1."subcategory_no" = 1
LEFT JOIN "post_subcategory" AS ps2
ON p."post_id" = ps2."post_id" AND ps2."subcategory_no" = 2
WHERE p."post_id" = $1;
  


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
