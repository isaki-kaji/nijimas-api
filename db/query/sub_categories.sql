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

-- name: GetUserUsedSubCategories :many
SELECT s.category_name, s.category_id
  FROM sub_categories s
  JOIN (
    SELECT ps.category_id, COUNT(*) AS post_count
      FROM posts p
      JOIN post_subcategories ps ON p.post_id = ps.post_id
     WHERE p.uid = $1
     GROUP BY ps.category_id
     ORDER BY post_count DESC
     LIMIT 20
  ) AS top_categories ON s.category_id = top_categories.category_id
ORDER BY top_categories.post_count DESC;
