// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: expense_summaries.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createExpenseSummary = `-- name: CreateExpenseSummary :one
INSERT INTO expense_summaries (uid, year, month, main_category, amount)
VALUES ($1, $2, $3, $4, $5)
RETURNING uid, year, month, main_category, amount
`

type CreateExpenseSummaryParams struct {
	Uid          string         `json:"uid"`
	Year         int32          `json:"year"`
	Month        int32          `json:"month"`
	MainCategory string         `json:"main_category"`
	Amount       pgtype.Numeric `json:"amount"`
}

func (q *Queries) CreateExpenseSummary(ctx context.Context, arg CreateExpenseSummaryParams) (ExpenseSummary, error) {
	row := q.db.QueryRow(ctx, createExpenseSummary,
		arg.Uid,
		arg.Year,
		arg.Month,
		arg.MainCategory,
		arg.Amount,
	)
	var i ExpenseSummary
	err := row.Scan(
		&i.Uid,
		&i.Year,
		&i.Month,
		&i.MainCategory,
		&i.Amount,
	)
	return i, err
}

const getExpenseSummariesByMonth = `-- name: GetExpenseSummariesByMonth :many
SELECT main_category, amount
FROM expense_summaries
WHERE uid = $1 AND year = $2 AND month = $3
`

type GetExpenseSummariesByMonthParams struct {
	Uid   string `json:"uid"`
	Year  int32  `json:"year"`
	Month int32  `json:"month"`
}

type GetExpenseSummariesByMonthRow struct {
	MainCategory string         `json:"main_category"`
	Amount       pgtype.Numeric `json:"amount"`
}

func (q *Queries) GetExpenseSummariesByMonth(ctx context.Context, arg GetExpenseSummariesByMonthParams) ([]GetExpenseSummariesByMonthRow, error) {
	rows, err := q.db.Query(ctx, getExpenseSummariesByMonth, arg.Uid, arg.Year, arg.Month)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetExpenseSummariesByMonthRow{}
	for rows.Next() {
		var i GetExpenseSummariesByMonthRow
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

const getExpenseSummary = `-- name: GetExpenseSummary :one
SELECT uid, year, month, main_category, amount FROM expense_summaries
WHERE uid = $1 AND year = $2 AND month = $3 AND main_category = $4
`

type GetExpenseSummaryParams struct {
	Uid          string `json:"uid"`
	Year         int32  `json:"year"`
	Month        int32  `json:"month"`
	MainCategory string `json:"main_category"`
}

func (q *Queries) GetExpenseSummary(ctx context.Context, arg GetExpenseSummaryParams) (ExpenseSummary, error) {
	row := q.db.QueryRow(ctx, getExpenseSummary,
		arg.Uid,
		arg.Year,
		arg.Month,
		arg.MainCategory,
	)
	var i ExpenseSummary
	err := row.Scan(
		&i.Uid,
		&i.Year,
		&i.Month,
		&i.MainCategory,
		&i.Amount,
	)
	return i, err
}

const updateExpenseSummary = `-- name: UpdateExpenseSummary :one
UPDATE expense_summaries
SET amount = $1
WHERE uid = $2 AND year = $3 AND month = $4 AND main_category = $5
RETURNING uid, year, month, main_category, amount
`

type UpdateExpenseSummaryParams struct {
	Amount       pgtype.Numeric `json:"amount"`
	Uid          string         `json:"uid"`
	Year         int32          `json:"year"`
	Month        int32          `json:"month"`
	MainCategory string         `json:"main_category"`
}

func (q *Queries) UpdateExpenseSummary(ctx context.Context, arg UpdateExpenseSummaryParams) (ExpenseSummary, error) {
	row := q.db.QueryRow(ctx, updateExpenseSummary,
		arg.Amount,
		arg.Uid,
		arg.Year,
		arg.Month,
		arg.MainCategory,
	)
	var i ExpenseSummary
	err := row.Scan(
		&i.Uid,
		&i.Year,
		&i.Month,
		&i.MainCategory,
		&i.Amount,
	)
	return i, err
}
