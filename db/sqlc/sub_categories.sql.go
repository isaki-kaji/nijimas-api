// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: sub_categories.sql

package db

import (
	"context"

	"github.com/google/uuid"
)

const createSubCategory = `-- name: CreateSubCategory :one
INSERT INTO sub_categories (
  category_id,
  category_name
) VALUES (
  $1, $2
) RETURNING category_id, category_name, created_at
`

type CreateSubCategoryParams struct {
	CategoryID   uuid.UUID `json:"category_id"`
	CategoryName string    `json:"category_name"`
}

func (q *Queries) CreateSubCategory(ctx context.Context, arg CreateSubCategoryParams) (SubCategory, error) {
	row := q.db.QueryRow(ctx, createSubCategory, arg.CategoryID, arg.CategoryName)
	var i SubCategory
	err := row.Scan(&i.CategoryID, &i.CategoryName, &i.CreatedAt)
	return i, err
}

const getSubCategoryByName = `-- name: GetSubCategoryByName :one
SELECT category_id, category_name, created_at FROM sub_categories
WHERE category_name = $1
`

func (q *Queries) GetSubCategoryByName(ctx context.Context, categoryName string) (SubCategory, error) {
	row := q.db.QueryRow(ctx, getSubCategoryByName, categoryName)
	var i SubCategory
	err := row.Scan(&i.CategoryID, &i.CategoryName, &i.CreatedAt)
	return i, err
}
