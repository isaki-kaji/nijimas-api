package testutil

import (
	"context"
	"math/rand"
	"strings"
	"time"

	"github.com/isaki-kaji/nijimas-api/configs"
	db "github.com/isaki-kaji/nijimas-api/db/sqlc"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func RandomString(n int) string {
	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var sb strings.Builder
	k := len(chars)
	for i := 0; i < n; i++ {
		c := chars[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

func RandomUid() string {
	return RandomString(28)
}

func RandomCountryCode() *string {
	countryCodes := []string{"JP", "US", "GB", "CA", "FR", "DE", "IT", "ES", "AU", "BR", "RU", "KR", "CN", "IN", "ID", "MX", "NL", "SA", "SE", "TR"}
	return &countryCodes[rand.Intn(len(countryCodes))]
}

func RandomMainCategory() string {
	mainCategories := []string{"食事", "趣味", "ファッション", "雑貨", "日用品", "旅行", "交際費", "交通費", "その他"}
	return mainCategories[rand.Intn(len(mainCategories))]
}

func RandomPublicTypeNo() string {
	publicTypeNos := []string{"1", "2", "3"}
	return publicTypeNos[rand.Intn(len(publicTypeNos))]
}

func RandomTime() time.Time {
	min := time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC)
	max := time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC)
	duration := max.Sub(min)
	randomDuration := time.Duration(rand.Int63n(int64(duration)))
	return min.Add(randomDuration)
}

func SetupDB() db.Repository {
	config, err := configs.LoadConfig("../../environment/development/")
	if err != nil {
		panic(err)
	}
	dbConfig, err := pgxpool.ParseConfig(config.DBSource)
	if err != nil {
		panic(err)
	}
	dbConfig.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
		_, err := conn.Exec(ctx, "SET TIME ZONE 'UTC'")
		if err != nil {
			panic(err)
		}
		return nil
	}
	connPool, err := pgxpool.NewWithConfig(context.Background(), dbConfig)
	if err != nil {
		panic(err)
	}
	return db.NewRepository(connPool)
}
