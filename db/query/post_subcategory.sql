-- name: CreatePostSubCategory :one
INSERT INTO "post_subcategory" (
  "post_id",
  "subcategory_no",
  "sub_category"
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: GetPostSubCategoryByPostId :many
SELECT * FROM "post_subcategory"
WHERE "post_id" = $1;
