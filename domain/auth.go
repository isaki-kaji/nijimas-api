package domain

import (
	"context"

	db "github.com/isaki-kaji/nijimas-api/db/sqlc"
)

type AuthService interface {
	SignupUser(ctx context.Context, arg db.CreateUserParams) (db.User, error)
}
