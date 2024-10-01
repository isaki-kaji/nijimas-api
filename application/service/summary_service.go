package service

import (
	"context"
	"sort"
	"time"

	"github.com/isaki-kaji/nijimas-api/apperror"
	db "github.com/isaki-kaji/nijimas-api/db/sqlc"
)

type SummaryService interface {
	GetMonthlySummary(ctx context.Context, uid string, year int, month int) (MonthlySummaryResponse, error)
}

func NewSummaryService(repository db.Repository) SummaryService {
	return &SummaryServiceImpl{repository: repository}
}

type SummaryServiceImpl struct {
	repository db.Repository
}

type MonthlySummaryResponse struct {
	Uid                string              `json:"uid"`
	Year               int                 `json:"year"`
	Month              int                 `json:"month"`
	ExpenseSummary     []CalculatedSummary `json:"expense_summary"`
	SubCategorySummary []CalculatedSummary `json:"subcategory_summary"`
	DailyCount         []int               `json:"daily_count"`
	DailyAmount        []int               `json:"daily_amount"`
}

type CalculatedSummary struct {
	CategoryName string  `json:"categoryName"`
	Amount       int     `json:"amount"`
	Percentage   float64 `json:"percentage"`
}

func (s *SummaryServiceImpl) GetMonthlySummary(ctx context.Context, uid string, year int, month int) (MonthlySummaryResponse, error) {
	startDate := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)
	endDate := startDate.AddDate(0, 1, 0)
	lastDayOfMonth := endDate.AddDate(0, 0, -1)
	daysInMonth := lastDayOfMonth.Day()

	getExpenseSummaryParam := db.GetExpenseSummaryByMonthParams{
		Uid:         uid,
		CreatedAt:   startDate,
		CreatedAt_2: endDate,
	}
	expenseSummary, err := s.repository.GetExpenseSummaryByMonth(ctx, getExpenseSummaryParam)
	if err != nil {
		err = apperror.GetDataFailed.Wrap(err, "failed to get expense summary")
		return MonthlySummaryResponse{}, err
	}
	calculatedExpenseSummary := processExpenseSummary(expenseSummary)

	getSubCategorySummaryParam := db.GetSubCategorySummaryByMonthParams{
		Uid:         uid,
		CreatedAt:   startDate,
		CreatedAt_2: endDate,
	}
	subCategorySummary, err := s.repository.GetSubCategorySummaryByMonth(ctx, getSubCategorySummaryParam)
	if err != nil {
		err = apperror.GetDataFailed.Wrap(err, "failed to get subcategory summary")
		return MonthlySummaryResponse{}, err
	}
	calculatedSubCategorySummary := processSubCategorySummary(subCategorySummary)

	getDailyActivitySummaryParam := db.GetDailyActivitySummaryByMonthParams{
		Uid:         uid,
		CreatedAt:   startDate,
		CreatedAt_2: endDate,
	}
	dailyActivitySummary, err := s.repository.GetDailyActivitySummaryByMonth(ctx, getDailyActivitySummaryParam)
	if err != nil {
		err = apperror.GetDataFailed.Wrap(err, "failed to get daily activity summary")
		return MonthlySummaryResponse{}, err
	}
	dailyCount, dailyAmount := generateDailyActivities(daysInMonth, dailyActivitySummary)

	monthlySummaryResponse := MonthlySummaryResponse{
		Uid:                uid,
		Year:               year,
		Month:              month,
		ExpenseSummary:     calculatedExpenseSummary,
		SubCategorySummary: calculatedSubCategorySummary,
		DailyCount:         dailyCount,
		DailyAmount:        dailyAmount,
	}

	return monthlySummaryResponse, nil
}

// 共通のロジック: TotalAmountを計算する
func calcTotalAmount[T any](summary []T, getAmount func(T) int) int {
	totalAmount := 0
	for _, row := range summary {
		totalAmount += getAmount(row)
	}
	return totalAmount
}

// 共通のロジック: Percentageを計算し、マップに格納する
func calcPercentage[T any](summary []T, getCategoryName func(T) string, getAmount func(T) int) map[string]CalculatedSummary {
	totalAmount := calcTotalAmount(summary, getAmount)
	percentageMap := make(map[string]CalculatedSummary)
	for _, row := range summary {
		percentage := float64(getAmount(row)) / float64(totalAmount) * 100
		percentageMap[getCategoryName(row)] = CalculatedSummary{
			CategoryName: getCategoryName(row),
			Amount:       getAmount(row),
			Percentage:   percentage,
		}
	}
	return percentageMap
}

// ソートされたCalculatedSummaryのスライスを返す
func buildSortedSummarySlice(summary map[string]CalculatedSummary) []CalculatedSummary {
	keys := make([]string, 0, len(summary))
	for key := range summary {
		keys = append(keys, key)
	}

	sort.Slice(keys, func(i, j int) bool {
		return summary[keys[i]].Amount > summary[keys[j]].Amount
	})

	sortedSummary := make([]CalculatedSummary, 0, len(summary))
	for _, key := range keys {
		sortedSummary = append(sortedSummary, summary[key])
	}

	return sortedSummary
}

// ExpenseSummaryの処理
func processExpenseSummary(expenseSummary []db.GetExpenseSummaryByMonthRow) []CalculatedSummary {
	// MainCategoryとAmountを取得するための関数を定義
	getCategoryName := func(row db.GetExpenseSummaryByMonthRow) string {
		return row.MainCategory
	}
	getAmount := func(row db.GetExpenseSummaryByMonthRow) int {
		return int(row.Amount)
	}

	// Percentageを計算し、ソートして返す
	percentageMap := calcPercentage(expenseSummary, getCategoryName, getAmount)
	return buildSortedSummarySlice(percentageMap)
}

// SubCategorySummaryの処理
func processSubCategorySummary(subCategorySummary []db.GetSubCategorySummaryByMonthRow) []CalculatedSummary {
	// CategoryNameとAmountを取得するための関数を定義
	getCategoryName := func(row db.GetSubCategorySummaryByMonthRow) string {
		return row.CategoryName
	}
	getAmount := func(row db.GetSubCategorySummaryByMonthRow) int {
		return int(row.Amount)
	}

	// Percentageを計算し、ソートして返す
	percentageMap := calcPercentage(subCategorySummary, getCategoryName, getAmount)
	return buildSortedSummarySlice(percentageMap)
}
func generateDailyActivities(daysInMonth int, dailyActivitySummary []db.GetDailyActivitySummaryByMonthRow) ([]int, []int) {

	// キャパシティを指定して初期化しているので、ゼロ値で初期化される
	dailyCount := make([]int, daysInMonth)
	dailyAmount := make([]int, daysInMonth)

	for _, row := range dailyActivitySummary {
		dailyCount[row.Date] = int(row.Count)
		dailyAmount[row.Date] = int(row.Amount)
	}

	return dailyCount, dailyAmount
}
