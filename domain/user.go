package domain

import (
	"context"

	db "github.com/isaki-kaji/nijimas-api/db/sqlc"
)

type UserService interface {
	CreateUser(ctx context.Context, arg CreateUserRequest) (db.User, error)
}

type CreateUserRequest struct {
	Uid         string `json:"uid" binding:"required"`
	Username    string `json:"username" binding:"required"`
	CountryCode string `json:"country_code" binding:"len=2"`
}
