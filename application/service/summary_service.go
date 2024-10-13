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
	GetMonthlySummary(ctx context.Context, uid string, year int, month int, timezone string) (MonthlySummaryResponse, error)
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
	Count        int     `json:"count"`
	Amount       int     `json:"amount"`
	Percentage   float64 `json:"percentage"`
}

func (s *SummaryServiceImpl) GetMonthlySummary(ctx context.Context, uid string, year int, month int, timezone string) (MonthlySummaryResponse, error) {
	loc, err := time.LoadLocation(timezone)
	if err != nil {
		return MonthlySummaryResponse{}, apperror.InvalidHeader.Wrap(err, "failed to load location")
	}
	startDate := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, loc)
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
			expenseSummaryChan <- calculatedSummaryResult{summary: []CalculatedSummary{}, err: err}
		}

		if len(expenseSummary) == 0 {
			expenseSummaryChan <- calculatedSummaryResult{summary: []CalculatedSummary{}, err: nil}
			return
		}

		calculatedExpenseSummary := processExpenseSummary(expenseSummary)
		expenseSummaryChan <- calculatedSummaryResult{summary: calculatedExpenseSummary, err: nil}
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
			subCategorySummaryChan <- calculatedSummaryResult{summary: []CalculatedSummary{}, err: err}
			return
		}

		if len(subCategorySummary) == 0 {
			subCategorySummaryChan <- calculatedSummaryResult{summary: []CalculatedSummary{}, err: nil}
			return
		}

		calculatedSubCategorySummary := processSubCategorySummary(subCategorySummary)
		subCategorySummaryChan <- calculatedSummaryResult{summary: calculatedSubCategorySummary, err: nil}
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
			dailyActivitySummaryChan <- dailyActivityResult{dailyCount: []int{}, dailyAmount: []int{}, err: err}
			return
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
func calcPercentage[T any](summary []T, getCategoryName func(T) string, getAmount func(T) int, getCount func(T) int) []CalculatedSummary {
	totalAmount := calcTotalAmount(summary, getAmount)
	calculatedSummaries := make([]CalculatedSummary, 0, len(summary))

	var totalPercentage float64

	// パーセンテージを計算し、小数点第1位で四捨五入
	for _, row := range summary {
		amount := getAmount(row)
		percentage := math.Round(float64(amount)/float64(totalAmount)*1000) / 10 // 小数点第1位で四捨五入
		totalPercentage += percentage

		calculatedSummaries = append(calculatedSummaries, CalculatedSummary{
			CategoryName: getCategoryName(row),
			Amount:       amount,
			Percentage:   percentage,
			Count:        getCount(row),
		})
	}

	// 合計を100%に調整する
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
		// 100%未満の場合、中央値の要素に加算する
		midIdx := findMidPercentageIdx(summaries)
		summaries[midIdx].Percentage += difference
	} else if difference < 0 {
		// 100%を超えている場合、中央値の要素から減算する
		midIdx := findMidPercentageIdx(summaries)
		summaries[midIdx].Percentage += difference // differenceは負の値なので減算される
	}

	// 負のパーセンテージがないか確認
	for i := range summaries {
		if summaries[i].Percentage < 0 {
			summaries[i].Percentage = 0 // 負のパーセンテージを0にする
		}
	}
}

// 一番割合が中央値に近い要素のインデックスを見つける
func findMidPercentageIdx(summaries []CalculatedSummary) int {
	n := len(summaries)
	if n == 0 {
		return 0
	}
	// 中央付近のインデックスを返す（ソートされている前提で）
	return n / 2
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

	getCount := func(ow db.GetExpenseSummaryByMonthRow) int {
		return 0
	}

	// Percentageを計算し、ソートされたスライスを返す
	return calcPercentage(expenseSummary, getCategoryName, getAmount, getCount)
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

	getCount := func(row db.GetSubCategorySummaryByMonthRow) int {
		return int(row.Count)
	}

	// Percentageを計算し、ソートされたスライスを返す
	return calcPercentage(subCategorySummary, getCategoryName, getAmount, getCount)
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
