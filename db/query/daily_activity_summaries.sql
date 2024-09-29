-- name: CreateDailyActivitySummary :one
INSERT INTO daily_activity_summaries (uid, year, month, day, number, amount)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: UpdateDailyActivitySummary :one
UPDATE daily_activity_summaries
SET number = $1, amount = $2
WHERE uid = $3 AND year = $4 AND month = $5 AND day = $6
RETURNING *;

-- name: GetDailyActivitySummary :one
SELECT * FROM daily_activity_summaries
WHERE uid = $1 AND year = $2 AND month = $3 AND day = $4;

-- name: GetDailyActivitySummariesByMonth :many
SELECT day, number, amount
FROM daily_activity_summaries
WHERE uid = $1 AND year = $2 AND month = $3
ORDER BY day ASC;
