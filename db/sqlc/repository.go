package db

import (
	"database/sql"
	"log"

	"github.com/isaki-kaji/nijimas-api/util"
	"go.uber.org/fx"
)

type Repository interface {
	Querier
}

type SQLRepository struct {
	*Queries
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &SQLRepository{
		db:      db,
		Queries: New(db),
	}
}

func NewDB(config *util.Config) (*sql.DB, error) {
	db, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
		return nil, err
	}
	return db, nil
}

var Module = fx.Options(
	fx.Provide(NewRepository),
	fx.Provide(NewDB),
)
