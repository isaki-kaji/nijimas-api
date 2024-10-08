-- name: GetExpenseSummaryByMonth :many
SELECT
  main_category,
  SUM(expense) AS amount
FROM posts
WHERE uid = $1
  AND (created_at AT TIME ZONE 'UTC' AT TIME ZONE 'Asia/Tokyo') >= $2
  AND (created_at AT TIME ZONE 'UTC' AT TIME ZONE 'Asia/Tokyo') < $3
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
  AND (p.created_at AT TIME ZONE 'UTC' AT TIME ZONE 'Asia/Tokyo') >= $2
  AND (p.created_at AT TIME ZONE 'UTC' AT TIME ZONE 'Asia/Tokyo') < $3
GROUP BY s.category_name;

-- name: GetDailyActivitySummaryByMonth :many
SELECT
  DATE_PART('day', (created_at AT TIME ZONE 'UTC' AT TIME ZONE 'Asia/Tokyo'))::int AS date,
  COUNT(*) AS count,
  SUM(expense) AS amount
FROM posts
WHERE uid = $1
  AND (created_at AT TIME ZONE 'UTC' AT TIME ZONE 'Asia/Tokyo') >= $2
  AND (created_at AT TIME ZONE 'UTC' AT TIME ZONE 'Asia/Tokyo') < $3
GROUP BY DATE_PART('day', (created_at AT TIME ZONE 'UTC' AT TIME ZONE 'Asia/Tokyo'))::int;