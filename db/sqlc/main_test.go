package db_test

import (
	"context"
	"log"
	"os"
	"testing"

	db "github.com/isaki-kaji/nijimas-api/db/sqlc"
	"github.com/isaki-kaji/nijimas-api/util"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var testRepository db.Repository

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../../environment/development/")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	dbConfig, err := pgxpool.ParseConfig(config.DBSource)
	if err != nil {
		log.Fatalf("Failed to parse db config: %v", err)
	}
	dbConfig.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
		_, err := conn.Exec(ctx, "SET TIME ZONE 'UTC'")
		return err
	}
	connPool, err := pgxpool.NewWithConfig(context.Background(), dbConfig)
	if err != nil {
		log.Fatalf("Failed to create connection pool: %v", err)
	}
	testRepository = db.NewRepository(connPool)
	os.Exit(m.Run())
}
