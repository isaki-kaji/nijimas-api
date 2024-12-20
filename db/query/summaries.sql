-- name: GetExpenseSummaryByMonth :many
SELECT
  main_category,
  SUM(expense) AS amount
FROM posts
WHERE uid = $1
  AND timezone('Asia/Tokyo', created_at) >= $2
  AND timezone('Asia/Tokyo', created_at) < $3
GROUP BY main_category;

-- name: GetSubCategorySummaryByMonth :many
SELECT
  s.category_name,
  COUNT(*) AS count,
  SUM(p.expense) AS amount
FROM posts p
JOIN post_subcategories ps ON p.post_id = ps.post_id
JOIN sub_categories s ON ps.category_id = s.category_id
WHERE p.uid = $1
  AND timezone('Asia/Tokyo', p.created_at) >= $2
  AND timezone('Asia/Tokyo', p.created_at) < $3
GROUP BY s.category_name;

-- name: GetDailyActivitySummaryByMonth :many
SELECT
  DATE_PART('day', timezone('Asia/Tokyo', created_at))::int AS date,
  COUNT(*) AS count,
  SUM(expense) AS amount
FROM posts
WHERE uid = $1
  AND timezone('Asia/Tokyo', created_at) >= $2
  AND timezone('Asia/Tokyo', created_at) < $3
GROUP BY DATE_PART('day', timezone('Asia/Tokyo', created_at))
ORDER BY date;

