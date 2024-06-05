-- name: CreatePost :one
INSERT INTO "post" (
  "post_id",
  "uid",
  "main_category",
  "post_text",
  "photo_url",
  "expense",
  "location",
  "public_type_no"
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8
) RETURNING *;  
  
-- name: GetPostById :one
SELECT
  p."post_id",
  u."uid",
  u."username",
  p."main_category",
  ps1."sub_category",
  ps2."sub_category",
  p."post_text",
  p."photo_url",
  p."expense",
  p."location",
  p."public_type_no",
  p."created_at"
FROM "post" AS p
JOIN "user" AS u ON p."uid" = u."uid"
LEFT JOIN "post_subcategory" AS ps1
ON p."post_id" = ps1."post_id" AND ps1."subcategory_no" = '1'
LEFT JOIN "post_subcategory" AS ps2
ON p."post_id" = ps2."post_id" AND ps2."subcategory_no" = '2'
WHERE p."post_id" = $1;
  
-- name: GetPostsByUid :many
SELECT
  p."post_id",
  u."uid",
  u."username",
  p."main_category",
  ps1."sub_category",
  ps2."sub_category",
  p."post_text",
  p."photo_url",
  p."expense",
  p."location",
  p."public_type_no",
  p."created_at",
  f."uid" IS NOT NULL AS "is_favorite"
FROM "post" AS p
JOIN "user" AS u ON p."uid" = u."uid"
LEFT JOIN "post_subcategory" AS ps1
ON p."post_id" = ps1."post_id" AND ps1."subcategory_no" = '1'
LEFT JOIN "post_subcategory" AS ps2
ON p."post_id" = ps2."post_id" AND ps2."subcategory_no" = '2'
LEFT JOIN "favorite" AS f
ON p."post_id" = f."post_id" AND f."uid" = $1
WHERE p."uid" = $1
ORDER BY p."created_at" DESC
LIMIT 50;

-- name: GetPostsByCategory :many
SELECT
  p."post_id",
  u."uid",
  u."username",
  p."main_category",
  ps1."sub_category",
  ps2."sub_category",
  p."post_text",
  p."photo_url",
  p."expense",
  p."location",
  p."public_type_no",
  p."created_at"
FROM "post" AS p
JOIN "user" AS u ON p."uid" = u."uid"
LEFT JOIN "post_subcategory" AS ps1
ON p."post_id" = ps1."post_id" AND ps1."subcategory_no" = '1'
LEFT JOIN "post_subcategory" AS ps2
ON p."post_id" = ps2."post_id" AND ps2."subcategory_no" = '2'
WHERE 
  (p."main_category" = $1 OR $1 IS NULL) AND
  (ps1."sub_category" = $2 OR $2 IS NULL) AND
  (ps2."sub_category" = $3 OR $3 IS NULL)
ORDER BY p."created_at" DESC
LIMIT 50;

-- name: GetPostsBySubCategory :many
SELECT
  p."post_id",
  u."uid",
  u."username",
  p."main_category",
  ps1."sub_category",
  ps2."sub_category",
  p."post_text",
  p."photo_url",
  p."expense",
  p."location",
  p."public_type_no",
  p."created_at"
FROM "post" AS p
JOIN "user" AS u ON p."uid" = u."uid"
LEFT JOIN "post_subcategory" AS ps1
ON p."post_id" = ps1."post_id" AND ps1."subcategory_no" = '1'
LEFT JOIN "post_subcategory" AS ps2
ON p."post_id" = ps2."post_id" AND ps2."subcategory_no" = '2'
WHERE 
  (ps1."sub_category" = $1 OR $1 IS NULL) AND
  (ps2."sub_category" = $2 OR $2 IS NULL)
ORDER BY p."created_at" DESC
LIMIT 50;

-- name: GetPostsByFollowing :many
SELECT 
  p."post_id",
  u."uid",
  u."username",
  p."main_category",
  ps1."sub_category",
  ps2."sub_category",
  p."post_text",
  p."photo_url",
  p."expense",
  p."location",
  p."public_type_no",
  p."created_at"
FROM "post" AS p
JOIN "user" AS u ON p."uid" = u."uid"
JOIN "follow_user" AS f ON f."follow_uid" = p."uid"
LEFT JOIN "post_subcategory" AS ps1
ON p."post_id" = ps1."post_id" AND ps1."subcategory_no" = '1'
LEFT JOIN "post_subcategory" AS ps2
ON p."post_id" = ps2."post_id" AND ps2."subcategory_no" = '2'
WHERE f."uid" = $1
ORDER BY p."created_at" DESC
LIMIT 50;

-- name: UpdatePost :one
UPDATE "post" SET
  "main_category" = COALESCE(sqlc.narg(main_category), "main_category"),
  "post_text" = COALESCE(sqlc.narg(post_text), "post_text"),
  "photo_url" = COALESCE(sqlc.narg(photo_url), "photo_url"),
  "expense" = COALESCE(sqlc.narg(expense), "expense"),
  "public_type_no" = COALESCE(sqlc.narg(public_type_no), "public_type_no")
WHERE "post_id" = sqlc.arg(post_id)
RETURNING *;
