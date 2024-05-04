package util

import (
	"math/rand"
	"strings"
)

func PointerOrNil[T comparable](value T) *T {
	var zero T
	if value == zero {
		return nil
	}
	return &value
}

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
