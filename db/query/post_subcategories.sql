-- name: CreatePostSubCategory :one
INSERT INTO post_subcategories (
  post_id,
  category_no,
  category_id
) VALUES (
  $1, $2, $3
) RETURNING *;
