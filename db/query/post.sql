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
  
-- name: GetPostsByUserId :many
SELECT
  p."post_id",
  p."main_category",
  ps1."sub_category",
  ps2."sub_category",
  p."post_text",
  p."photo_url",
  p."location",
  p."public_type_no"
FROM "post" AS p
LEFT JOIN "post_subcategory" AS ps1
ON p."post_id" = ps1."post_id" AND ps1."subcategory_no" = 1
LEFT JOIN "post_subcategory" AS ps2
ON p."post_id" = ps2."post_id" AND ps2."subcategory_no" = 2
WHERE p."user_id" = $1
ORDER BY p."created_at" DESC
LIMIT 50;

-- name: GetPostsByCategory :many
SELECT 
  p."post_id",
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
WHERE 
  (p."main_category" = $1 OR $1 IS NULL) AND
  (ps1."sub_category" = $2 OR $2 IS NULL) AND
  (ps2."sub_category" = $3 OR $3 IS NULL)
ORDER BY p."created_at" DESC
LIMIT 50;

-- name: GetPostsBySubCategory :many
SELECT 
  p."post_id",
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
WHERE 
  (ps1."sub_category" = $1 OR $1 IS NULL) AND
  (ps2."sub_category" = $2 OR $2 IS NULL)
ORDER BY p."created_at" DESC
LIMIT 50;

-- name: GetPostsByFollowing :many
SELECT 
  p."post_id",
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
JOIN "follow_user" AS f ON f."follow_user_id" = p."user_id"
LEFT JOIN "post_subcategory" AS ps1
ON p."post_id" = ps1."post_id" AND ps1."subcategory_no" = 1
LEFT JOIN "post_subcategory" AS ps2
ON p."post_id" = ps2."post_id" AND ps2."subcategory_no" = 2
WHERE f."user_id" = $1
ORDER BY p."created_at" DESC
LIMIT 50;

-- name: UpdatePost :one
UPDATE "post" SET
  "main_category" = COALESCE(sqlc.narg(main_category), "main_category"),
  "post_text" = COALESCE(sqlc.narg(post_text), "post_text"),
  "photo_url" = COALESCE(sqlc.narg(photo_url), "photo_url"),
  "public_type_no" = COALESCE(sqlc.narg(public_type_no), "public_type_no")
WHERE "post_id" = sqlc.arg(post_id)
RETURNING *;