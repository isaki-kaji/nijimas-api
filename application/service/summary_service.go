package service

import (
	"context"
	"math"
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

	var expenseSummary []CalculatedSummary
	var subCategorySummary []CalculatedSummary
	var dailyCount []int
	var dailyAmount []int

	type calculatedSummaryResult struct {
		summary []CalculatedSummary
		err     error
	}

	type dailyActivityResult struct {
		dailyCount  []int
		dailyAmount []int
		err         error
	}
	expenseSummaryChan := make(chan calculatedSummaryResult)
	subCategorySummaryChan := make(chan calculatedSummaryResult)
	dailyActivitySummaryChan := make(chan dailyActivityResult)

	defer close(expenseSummaryChan)
	defer close(subCategorySummaryChan)
	defer close(dailyActivitySummaryChan)

	go func() {
		getExpenseSummaryParam := db.GetExpenseSummaryByMonthParams{
			Uid:         uid,
			CreatedAt:   startDate,
			CreatedAt_2: endDate,
		}
		expenseSummary, err := s.repository.GetExpenseSummaryByMonth(ctx, getExpenseSummaryParam)
		if err != nil {
			err = apperror.GetDataFailed.Wrap(err, "failed to get expense summary")
		}
		calculatedExpenseSummary := processExpenseSummary(expenseSummary)
		expenseSummaryChan <- calculatedSummaryResult{summary: calculatedExpenseSummary, err: err}
	}()

	go func() {
		getSubCategorySummaryParam := db.GetSubCategorySummaryByMonthParams{
			Uid:         uid,
			CreatedAt:   startDate,
			CreatedAt_2: endDate,
		}
		subCategorySummary, err := s.repository.GetSubCategorySummaryByMonth(ctx, getSubCategorySummaryParam)
		if err != nil {
			err = apperror.GetDataFailed.Wrap(err, "failed to get subcategory summary")
		}
		calculatedSubCategorySummary := processSubCategorySummary(subCategorySummary)
		subCategorySummaryChan <- calculatedSummaryResult{summary: calculatedSubCategorySummary, err: err}
	}()

	go func() {
		getDailyActivitySummaryParam := db.GetDailyActivitySummaryByMonthParams{
			Uid:         uid,
			CreatedAt:   startDate,
			CreatedAt_2: endDate,
		}
		dailyActivitySummary, err := s.repository.GetDailyActivitySummaryByMonth(ctx, getDailyActivitySummaryParam)
		if err != nil {
			err = apperror.GetDataFailed.Wrap(err, "failed to get daily activity summary")
		}
		dailyCount, dailyAmount := generateDailyActivities(daysInMonth, dailyActivitySummary)
		dailyActivitySummaryChan <- dailyActivityResult{dailyCount: dailyCount, dailyAmount: dailyAmount, err: err}
	}()

	for i := 0; i < 3; i++ {
		select {
		case er := <-expenseSummaryChan:
			if er.err != nil {
				return MonthlySummaryResponse{}, er.err
			}
			expenseSummary = er.summary
		case sr := <-subCategorySummaryChan:
			if sr.err != nil {
				return MonthlySummaryResponse{}, sr.err
			}
			subCategorySummary = sr.summary
		case dr := <-dailyActivitySummaryChan:
			if dr.err != nil {
				return MonthlySummaryResponse{}, dr.err
			}
			dailyCount = dr.dailyCount
			dailyAmount = dr.dailyAmount
		}
	}

	monthlySummaryResponse := MonthlySummaryResponse{
		Uid:                uid,
		Year:               year,
		Month:              month,
		ExpenseSummary:     expenseSummary,
		SubCategorySummary: subCategorySummary,
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

// 共通のロジック: パーセンテージを計算し、スライスに格納する
func calcPercentage[T any](summary []T, getCategoryName func(T) string, getAmount func(T) int) []CalculatedSummary {
	totalAmount := calcTotalAmount(summary, getAmount)
	calculatedSummaries := make([]CalculatedSummary, 0, len(summary))

	var totalPercentage float64

	// パーセンテージを計算し、小数点第1位で四捨五入
	for _, row := range summary {
		amount := getAmount(row)
		percentage := math.Floor(float64(amount)/float64(totalAmount)*1000) / 10 // 小数点第1位で切り捨て
		totalPercentage += percentage

		calculatedSummaries = append(calculatedSummaries, CalculatedSummary{
			CategoryName: getCategoryName(row),
			Amount:       amount,
			Percentage:   percentage,
		})
	}

	// 小数点第1位までに丸めて差分を計算
	adjustPercentageToHundred(calculatedSummaries, totalPercentage)

	// Amountでソート
	sort.Slice(calculatedSummaries, func(i, j int) bool {
		return calculatedSummaries[i].Amount > calculatedSummaries[j].Amount
	})

	return calculatedSummaries
}

// 合計を100%に調整する
func adjustPercentageToHundred(summaries []CalculatedSummary, totalPercentage float64) {
	difference := math.Round((100-totalPercentage)*10) / 10 // 小数点第1位で四捨五入して差分を計算

	if difference > 0 {
		// 100%未満の場合、一番割合が高い要素に加算
		highestIdx := findHighestPercentageIdx(summaries)
		summaries[highestIdx].Percentage += difference
	} else if difference < 0 {
		// 100%を超えている場合、一番割合が低い要素から減算
		lowestIdx := findLowestPercentageIdx(summaries)
		summaries[lowestIdx].Percentage += difference // differenceは負の値なので減算される
	}
}

// // パーセントを調整しても順序が変わらないようにしている ////
// 最も割合が高い要素を見つける
func findHighestPercentageIdx(summaries []CalculatedSummary) int {
	highestIdx := 0
	for i := range summaries {
		if summaries[i].Percentage > summaries[highestIdx].Percentage {
			highestIdx = i
		}
	}
	return highestIdx
}

// 最も割合が低い要素を見つける
func findLowestPercentageIdx(summaries []CalculatedSummary) int {
	lowestIdx := 0
	for i := range summaries {
		if summaries[i].Percentage < summaries[lowestIdx].Percentage {
			lowestIdx = i
		}
	}
	return lowestIdx
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

	// Percentageを計算し、ソートされたスライスを返す
	return calcPercentage(expenseSummary, getCategoryName, getAmount)
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

	// Percentageを計算し、ソートされたスライスを返す
	return calcPercentage(subCategorySummary, getCategoryName, getAmount)
}

// 日別のアクティビティデータを生成する
func generateDailyActivities(daysInMonth int, dailyActivitySummary []db.GetDailyActivitySummaryByMonthRow) ([]int, []int) {
	// キャパシティを指定して初期化しているので、ゼロ値で初期化される
	dailyCount := make([]int, daysInMonth)
	dailyAmount := make([]int, daysInMonth)

	for _, row := range dailyActivitySummary {
		dailyCount[row.Date-1] = int(row.Count)
		dailyAmount[row.Date-1] = int(row.Amount)
	}

	return dailyCount, dailyAmount
}
