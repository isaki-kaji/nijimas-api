-- name: CreatePost :one
INSERT INTO posts (
  post_id,
  uid,
  main_category,
  post_text,
  photo_url,
  expense,
  location,
  public_type_no
) VALUES (
  $1,
  $2,
  $3,
  CASE WHEN sqlc.narg(post_text)::text = '' THEN NULL ELSE sqlc.narg(post_text)::text END,
  CASE WHEN sqlc.narg(photo_url)::text = '' THEN NULL ELSE sqlc.narg(photo_url)::text END,
  $4,
  CASE WHEN sqlc.narg(location)::text = '' THEN NULL ELSE sqlc.narg(location)::text END,
  $5
) RETURNING *;

-- name: GetOwnPosts :many
SELECT
  p.post_id,
  u.uid,
  u.username,
  u.profile_image_url,
  p.main_category,
  COALESCE(sc.sub_category1, '')::text AS subCategory1,
  COALESCE(sc.sub_category2, '')::text AS subCategory2,
  p.post_text,
  p.photo_url,
  p.expense,
  p.location,
  CASE WHEN f.uid IS NOT NULL THEN TRUE ELSE FALSE END AS is_favorite,
  p.public_type_no,
  p.created_at
FROM posts p
JOIN users u ON p.uid = u.uid
LEFT JOIN (
  SELECT
    ps.post_id,
    MAX(CASE WHEN ps.category_no = '1' THEN s.category_name ELSE NULL END) AS sub_category1,
    MAX(CASE WHEN ps.category_no = '2' THEN s.category_name ELSE NULL END) AS sub_category2
  FROM post_subcategories ps
  JOIN sub_categories s ON ps.category_id = s.category_id
  GROUP BY ps.post_id
) sc ON p.post_id = sc.post_id
LEFT JOIN favorites f
  ON p.post_id = f.post_id AND f.uid = $1
WHERE p.uid = $1
ORDER BY p.post_id DESC
LIMIT 50;

-- name: GetTimelinePosts :many
SELECT
  p.post_id,
  u.uid,
  u.username,
  u.profile_image_url,
  p.main_category,
  COALESCE(sc.sub_category1, '')::text AS subCategory1,
  COALESCE(sc.sub_category2, '')::text AS subCategory2,
  p.post_text,
  p.photo_url,
  p.expense,
  p.location,
  CASE WHEN f.uid IS NOT NULL THEN TRUE ELSE FALSE END AS is_favorite,
  p.public_type_no,
  p.created_at
FROM posts p
JOIN users u ON p.uid = u.uid
LEFT JOIN (
  SELECT
    ps.post_id,
    MAX(CASE WHEN ps.category_no = '1' THEN s.category_name ELSE NULL END) AS sub_category1,
    MAX(CASE WHEN ps.category_no = '2' THEN s.category_name ELSE NULL END) AS sub_category2
  FROM post_subcategories ps
  JOIN sub_categories s ON ps.category_id = s.category_id
  GROUP BY ps.post_id
) sc ON p.post_id = sc.post_id
LEFT JOIN favorites f
  ON p.post_id = f.post_id AND f.uid = $1
WHERE 
  p.uid = $1
  OR (
    p.public_type_no IN ('0', '1') 
    AND EXISTS (
      SELECT 1 
      FROM follows f
      WHERE f.uid = $1 AND f.following_uid = p.uid
    )
  )
ORDER BY p.post_id DESC
LIMIT 50;


-- name: GetPostById :one
SELECT
  p.post_id,
  u.uid,
  u.username,
  u.profile_image_url,
  p.main_category,
  sc.sub_category1 AS subCategory1,
  sc.sub_category2 AS subCategory2,
  p.post_text,
  p.photo_url,
  p.expense,
  p.location,
  p.public_type_no,
  p.created_at
FROM posts p
JOIN users u ON p.uid = u.uid
LEFT JOIN (
  SELECT
    ps.post_id,
    MAX(CASE WHEN ps.category_no = '1' THEN s.category_name ELSE NULL END)::string AS sub_category1,
    MAX(CASE WHEN ps.category_no = '2' THEN s.category_name ELSE NULL END)::string AS sub_category2
  FROM post_subcategories ps
  JOIN sub_categories s ON ps.category_id = s.category_id
  GROUP BY ps.post_id
) sc ON p.post_id = sc.post_id
WHERE p.post_id = $1;

-- name: UpdatePost :one
UPDATE posts SET
  main_category = COALESCE(sqlc.narg(main_category), main_category),
  post_text = COALESCE(sqlc.narg(post_text), post_text),
  photo_url = COALESCE(sqlc.narg(photo_url), photo_url),
  expense = COALESCE(sqlc.narg(expense), expense),
  public_type_no = COALESCE(sqlc.narg(public_type_no), public_type_no)
WHERE post_id = sqlc.arg(post_id)
RETURNING *;

-- name: DeletePost :exec
DELETE FROM posts
WHERE post_id = $1;

-- name: GetPostsCount :one
SELECT COUNT(*) AS count FROM posts WHERE uid = $1;