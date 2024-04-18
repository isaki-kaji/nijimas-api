-- name: CreateSubCategory :one
INSERT INTO "sub_category" (
  "category_name"
) VALUES (
  $1
) RETURNING *;

-- name: GetSubCategory :one
SELECT * FROM "sub_category"
WHERE "category_name" = $1;
