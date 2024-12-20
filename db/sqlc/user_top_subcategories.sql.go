// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: user_top_subcategories.sql

package db

import (
	"context"

	"github.com/google/uuid"
)

const createUserTopSubCategories = `-- name: CreateUserTopSubCategories :many
INSERT INTO user_top_subcategories (uid, category_no, category_id)
VALUES
  ($1, '1', $2),
  ($1, '2', $3),
  ($1, '3', $4),
  ($1, '4', $5)
ON CONFLICT (uid, category_no)
DO UPDATE SET 
    category_id = EXCLUDED.category_id
WHERE user_top_subcategories.category_id IS DISTINCT FROM EXCLUDED.category_id
RETURNING uid, category_no, category_id, created_at
`

type CreateUserTopSubCategoriesParams struct {
	Uid          string    `json:"uid"`
	CategoryID   uuid.UUID `json:"category_id"`
	CategoryID_2 uuid.UUID `json:"category_id_2"`
	CategoryID_3 uuid.UUID `json:"category_id_3"`
	CategoryID_4 uuid.UUID `json:"category_id_4"`
}

func (q *Queries) CreateUserTopSubCategories(ctx context.Context, arg CreateUserTopSubCategoriesParams) ([]UserTopSubcategory, error) {
	rows, err := q.db.Query(ctx, createUserTopSubCategories,
		arg.Uid,
		arg.CategoryID,
		arg.CategoryID_2,
		arg.CategoryID_3,
		arg.CategoryID_4,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []UserTopSubcategory{}
	for rows.Next() {
		var i UserTopSubcategory
		if err := rows.Scan(
			&i.Uid,
			&i.CategoryNo,
			&i.CategoryID,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
