-- name: GetPostsByUid :many
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
WHERE p.uid = $2 
  AND (
    p.public_type_no = '0'
    OR (
      p.public_type_no = '1'
      AND EXISTS (
        SELECT 1 
        FROM follows 
        WHERE uid = $1 AND following_uid = p.uid
      )
    )
  )
ORDER BY p.post_id DESC
LIMIT 50;

-- name: GetPostsByMainCategory :many
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
WHERE p.main_category = $2 AND 
  (
    p.public_type_no = '0'
    OR (
      p.public_type_no = '1'
      AND EXISTS (
        SELECT 1 
        FROM follows 
        WHERE uid = $1 AND following_uid = p.uid
      )
    )
  )
ORDER BY p.post_id DESC
LIMIT 50;

-- name: GetPostsBySubCategory :many
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
  p.created_at
FROM posts p
JOIN users u ON p.uid = u.uid
JOIN (
  SELECT
    ps.post_id,
    MAX(CASE WHEN ps.category_no = '1' THEN s.category_name ELSE NULL END) AS sub_category1,
    MAX(CASE WHEN ps.category_no = '2' THEN s.category_name ELSE NULL END) AS sub_category2
  FROM post_subcategories ps
  JOIN sub_categories s ON ps.category_id = s.category_id
  GROUP BY ps.post_id
  HAVING 
    MAX(CASE WHEN ps.category_no = '1' THEN s.category_name ELSE NULL END) = sqlc.arg(category_name) OR 
    MAX(CASE WHEN ps.category_no = '2' THEN s.category_name ELSE NULL END) = sqlc.arg(category_name) 
) sc ON p.post_id = sc.post_id
LEFT JOIN favorites f
  ON p.post_id = f.post_id AND f.uid = $1
WHERE p.public_type_no = '0'
    OR (
      p.public_type_no = '1'
      AND EXISTS (
        SELECT 1 
        FROM follows 
        WHERE uid = $1 AND following_uid = p.uid
      )
    )
ORDER BY p.post_id DESC
LIMIT 50;

-- name: GetPostsByMainCategoryAndSubCategory :many
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
  p.created_at
FROM posts p
JOIN users u ON p.uid = u.uid
JOIN (
  SELECT
    ps.post_id,
    MAX(CASE WHEN ps.category_no = '1' THEN s.category_name ELSE NULL END) AS sub_category1,
    MAX(CASE WHEN ps.category_no = '2' THEN s.category_name ELSE NULL END) AS sub_category2
  FROM post_subcategories ps
  JOIN sub_categories s ON ps.category_id = s.category_id
  GROUP BY ps.post_id
  HAVING 
    MAX(CASE WHEN ps.category_no = '1' THEN s.category_name ELSE NULL END) = sqlc.arg(sub_category) OR 
    MAX(CASE WHEN ps.category_no = '2' THEN s.category_name ELSE NULL END) = sqlc.arg(sub_category) 
) sc ON p.post_id = sc.post_id
LEFT JOIN favorites f
  ON p.post_id = f.post_id AND f.uid = $1
WHERE p.main_category = $2 AND 
  (
    p.public_type_no = '0'
    OR (
      p.public_type_no = '1'
      AND EXISTS (
        SELECT 1 
        FROM follows 
        WHERE uid = $1 AND following_uid = p.uid
      )
    )
  )
ORDER BY p.post_id DESC
LIMIT 50;
