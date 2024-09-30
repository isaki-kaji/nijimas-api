-- name: CreatePostSubCategory :one
INSERT INTO post_subcategories (
  post_id,
  category_no,
  category_id
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: DeletePostSubCategory :exec
DELETE FROM post_subcategories
WHERE post_id = $1 AND category_no = $2;