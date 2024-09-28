package util

import (
	"github.com/bwmarrin/snowflake"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

func ToPointerOrNil[T comparable](value T) *T {
	var zero T
	if value == zero {
		return nil
	}
	return &value
}

func ToNumeric(expense string) (pgtype.Numeric, error) {
	var num pgtype.Numeric
	if err := num.Scan(expense); err != nil {
		return pgtype.Numeric{}, err
	}
	return num, nil
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
