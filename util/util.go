package util

import (
	"github.com/bwmarrin/snowflake"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

func ToPointerOrNil[T comparable](value T) *T {
	var zero T
	if value == zero {
		return nil
	}
	return &value
}

func ToDecimal(expense string) (decimal.Decimal, error) {
	// decimal.Decimal型の変数を作成
	dec, err := decimal.NewFromString(expense)
	if err != nil {
		return decimal.Decimal{}, err
	}
	return dec, nil
}

func StringPointerToString(p *string) string {
	if p == nil {
		return ""
	}
	return *p
}

func GenerateUUID() (uuid.UUID, error) {
	uuid, err := uuid.NewV7()
	if err != nil {
		return uuid, err
	}
	return uuid, nil
}

func GenerateSnowflakeID() (string, error) {
	node, err := snowflake.NewNode(1)
	if err != nil {
		return "", err
	}
	id := node.Generate().String()
	return id, nil
}
