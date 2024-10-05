package db

import (
	"context"

	"github.com/isaki-kaji/nijimas-api/configs"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/fx"
)

type Repository interface {
	Querier
	CreatePostTx(ctx context.Context, param CreatePostTxParam) (Post, error)
}

type SQLRepository struct {
	connPool *pgxpool.Pool
	*Queries
}

func NewRepository(connPool *pgxpool.Pool) Repository {
	return &SQLRepository{
		connPool: connPool,
		Queries:  New(connPool),
	}
}

func NewPool(config *configs.Config) (*pgxpool.Pool, error) {
	dbConfig, err := pgxpool.ParseConfig(config.DBSource)
	if err != nil {
		return nil, err
	}
	dbConfig.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
		_, err := conn.Exec(ctx, "SET TIME ZONE 'UTC'")
		return err
	}
	return pgxpool.NewWithConfig(context.Background(), dbConfig)
}

var Module = fx.Options(
	fx.Provide(NewRepository),
	fx.Provide(NewPool),
)
