// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: summaries.sql

package db

import (
	"context"
	"time"
)

const getDailyActivitySummaryByMonth = `-- name: GetDailyActivitySummaryByMonth :many
SELECT
  DATE_PART('day', timezone('Asia/Tokyo', created_at))::int AS date,
  COUNT(*) AS count,
  SUM(expense) AS amount
FROM posts
WHERE uid = $1
  AND timezone('Asia/Tokyo', created_at) >= $2
  AND timezone('Asia/Tokyo', created_at) < $3
GROUP BY DATE_PART('day', timezone('Asia/Tokyo', created_at))
ORDER BY date
`

type GetDailyActivitySummaryByMonthParams struct {
	Uid         string    `json:"uid"`
	CreatedAt   time.Time `json:"created_at"`
	CreatedAt_2 time.Time `json:"created_at_2"`
}

type GetDailyActivitySummaryByMonthRow struct {
	Date   int32 `json:"date"`
	Count  int64 `json:"count"`
	Amount int64 `json:"amount"`
}

func (q *Queries) GetDailyActivitySummaryByMonth(ctx context.Context, arg GetDailyActivitySummaryByMonthParams) ([]GetDailyActivitySummaryByMonthRow, error) {
	rows, err := q.db.Query(ctx, getDailyActivitySummaryByMonth, arg.Uid, arg.CreatedAt, arg.CreatedAt_2)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetDailyActivitySummaryByMonthRow{}
	for rows.Next() {
		var i GetDailyActivitySummaryByMonthRow
		if err := rows.Scan(&i.Date, &i.Count, &i.Amount); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getExpenseSummaryByMonth = `-- name: GetExpenseSummaryByMonth :many
SELECT
  main_category,
  SUM(expense) AS amount
FROM posts
WHERE uid = $1
  AND timezone('Asia/Tokyo', created_at) >= $2
  AND timezone('Asia/Tokyo', created_at) < $3
GROUP BY main_category
`

type GetExpenseSummaryByMonthParams struct {
	Uid         string    `json:"uid"`
	CreatedAt   time.Time `json:"created_at"`
	CreatedAt_2 time.Time `json:"created_at_2"`
}

type GetExpenseSummaryByMonthRow struct {
	MainCategory string `json:"main_category"`
	Amount       int64  `json:"amount"`
}

func (q *Queries) GetExpenseSummaryByMonth(ctx context.Context, arg GetExpenseSummaryByMonthParams) ([]GetExpenseSummaryByMonthRow, error) {
	rows, err := q.db.Query(ctx, getExpenseSummaryByMonth, arg.Uid, arg.CreatedAt, arg.CreatedAt_2)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetExpenseSummaryByMonthRow{}
	for rows.Next() {
		var i GetExpenseSummaryByMonthRow
		if err := rows.Scan(&i.MainCategory, &i.Amount); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getSubCategorySummaryByMonth = `-- name: GetSubCategorySummaryByMonth :many
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
GROUP BY s.category_name
`

type GetSubCategorySummaryByMonthParams struct {
	Uid         string    `json:"uid"`
	CreatedAt   time.Time `json:"created_at"`
	CreatedAt_2 time.Time `json:"created_at_2"`
}

type GetSubCategorySummaryByMonthRow struct {
	CategoryName string `json:"category_name"`
	Count        int64  `json:"count"`
	Amount       int64  `json:"amount"`
}

func (q *Queries) GetSubCategorySummaryByMonth(ctx context.Context, arg GetSubCategorySummaryByMonthParams) ([]GetSubCategorySummaryByMonthRow, error) {
	rows, err := q.db.Query(ctx, getSubCategorySummaryByMonth, arg.Uid, arg.CreatedAt, arg.CreatedAt_2)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetSubCategorySummaryByMonthRow{}
	for rows.Next() {
		var i GetSubCategorySummaryByMonthRow
		if err := rows.Scan(&i.CategoryName, &i.Count, &i.Amount); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
