-- name: CreateMainCategory :one
INSERT INTO "main_category" (
 "category_name"
) VALUES (
 $1
) RETURNING *;

-- name: GetMainCategory :one
SELECT * FROM "main_category"
WHERE "category_name" = $1;

-- name: GetMainCategories :many
SELECT * FROM "main_category";
