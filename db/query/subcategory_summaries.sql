-- name: CreateSubCategorySummary :one
INSERT INTO subcategory_summaries (uid, year, month, category_id, amount)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: UpdateSubCategorySummary :one
UPDATE subcategory_summaries
SET amount = $1
WHERE uid = $2 AND year = $3 AND month = $4 AND category_id = $5
RETURNING *;

-- name: GetSubCategorySummary :one
SELECT * FROM subcategory_summaries
WHERE uid = $1 AND year = $2 AND month = $3 AND category_id = $4;

-- name: GetSubCategorySummariesByMonth :many
SELECT category_id, amount
FROM subcategory_summaries
WHERE uid = $1 AND year = $2 AND month = $3;
