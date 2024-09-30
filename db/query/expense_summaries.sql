-- name: CreateExpenseSummary :one
INSERT INTO expense_summaries (uid, year, month, main_category, amount)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: UpdateExpenseSummary :one
UPDATE expense_summaries
SET amount = $1
WHERE uid = $2 AND year = $3 AND month = $4 AND main_category = $5
RETURNING *;

-- name: GetExpenseSummary :one
SELECT * FROM expense_summaries
WHERE uid = $1 AND year = $2 AND month = $3 AND main_category = $4;

-- name: GetExpenseSummariesByMonth :many
SELECT main_category, amount
FROM expense_summaries
WHERE uid = $1 AND year = $2 AND month = $3;

-- name: DeleteExpenseSummary :exec
DELETE FROM expense_summaries
WHERE uid = $1 AND year = $2 AND month = $3 AND main_category = $4;