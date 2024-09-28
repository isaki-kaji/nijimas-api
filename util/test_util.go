package util

import (
	"math/rand"
	"strings"
	"time"
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
