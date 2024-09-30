-- name: CreateSubCategory :one
INSERT INTO sub_categories (
  category_id,
  category_name
) VALUES (
  $1, $2
) RETURNING *;

-- name: GetSubCategoryByName :one
SELECT * FROM sub_categories
WHERE category_name = $1;

-- name: DeleteSubCategory :exec
DELETE FROM sub_categories
WHERE category_id = $1;