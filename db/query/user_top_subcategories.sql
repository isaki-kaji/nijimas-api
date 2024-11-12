-- name: CreateUserTopSubCategories :many
INSERT INTO user_top_subcategories (uid, category_no, category_id)
VALUES
  ($1, '1', $2),
  ($1, '2', $3),
  ($1, '3', $4),
  ($1, '4', $5)
ON CONFLICT (uid, category_no)
DO UPDATE SET 
    category_id = EXCLUDED.category_id
WHERE user_top_subcategories.category_id IS DISTINCT FROM EXCLUDED.category_id
RETURNING *;
