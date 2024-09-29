package db

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/isaki-kaji/nijimas-api/util"
	"github.com/jackc/pgx/v5"
	"github.com/shopspring/decimal"
)

var MaxAmount = decimal.NewFromInt(99999999)
var MinAmount = decimal.NewFromInt(0)

type CreatePostTxParam struct {
	PostID       uuid.UUID
	Uid          string
	MainCategory string
	SubCategory1 string
	SubCategory2 string
	PostText     *string
	PhotoUrl     *string
	Expense      string
	Location     *string
	PublicTypeNo string
	CreatedAt    time.Time
}

func (r *SQLRepository) CreatePostTx(ctx context.Context, param CreatePostTxParam) (Post, error) {
	tx, err := r.connPool.Begin(ctx)
	if err != nil {
		return Post{}, err
	}
	defer func() {
		// 正常にコミットされた場合は、ロールバックしない
		if err != nil {
			tx.Rollback(ctx)
		}
	}()
	qtx := r.WithTx(tx)

	decimalExpense, err := util.ToDecimal(param.Expense)
	if err != nil {
		return Post{}, err
	}

	dbParam := CreatePostParams{
		PostID:       param.PostID,
		Uid:          param.Uid,
		MainCategory: param.MainCategory,
		PostText:     param.PostText,
		PhotoUrl:     param.PhotoUrl,
		Expense:      decimalExpense,
		Location:     param.Location,
		PublicTypeNo: param.PublicTypeNo,
	}

	post, err := qtx.CreatePost(ctx, dbParam)
	if err != nil {
		return Post{}, err
	}

	//categoryIdをバッファ付きチャネルに入れたら良さそう
	categoryId1, err := handleSubCategory(ctx, param.PostID, param.SubCategory1, "1", qtx)
	if err != nil {
		return Post{}, err
	}
	categoryId2, err := handleSubCategory(ctx, param.PostID, param.SubCategory2, "2", qtx)
	if err != nil {
		return Post{}, err
	}

	CalcSummaryParam := CalcSummaryParam{
		Uid:          param.Uid,
		MainCategory: param.MainCategory,
		Expense:      decimalExpense,
		year:         int32(param.CreatedAt.Year()),
		month:        int32(param.CreatedAt.Month()),
		day:          int32(param.CreatedAt.Day()),
	}

	err = calcExpenseSummary(ctx, CalcSummaryParam, qtx)
	if err != nil {
		return Post{}, err
	}

	//for-selectが使えそう
	err = calcSubcategorySummary(ctx, CalcSummaryParam, categoryId1, qtx)
	if err != nil {
		return Post{}, err
	}
	err = calcSubcategorySummary(ctx, CalcSummaryParam, categoryId2, qtx)
	if err != nil {
		return Post{}, err
	}

	err = calcDailyActivitySummary(ctx, CalcSummaryParam, qtx)
	if err != nil {
		return Post{}, err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return Post{}, err
	}

	return post, nil
}

func handleSubCategory(ctx context.Context, postID uuid.UUID, categoryName string, categoryNo string, qtx *Queries) (uuid.UUID, error) {
	if categoryName == "" {
		return uuid.Nil, nil
	}

	categoryId, err := registerSubCategory(ctx, categoryName, qtx)
	if err != nil {
		return uuid.Nil, err
	}

	err = registerPostSubCategory(ctx, postID, categoryId, categoryNo, qtx)
	if err != nil {
		return uuid.Nil, err
	}
	return categoryId, nil
}

func registerSubCategory(ctx context.Context, categoryName string, qtx *Queries) (uuid.UUID, error) {
	subCategory, err := qtx.GetSubCategoryByName(ctx, categoryName)
	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			return uuid.Nil, err
		}

		categoryId, err := util.GenerateUUID()
		if err != nil {
			return uuid.Nil, err
		}

		createSubCategoryParam := CreateSubCategoryParams{
			CategoryID:   categoryId,
			CategoryName: categoryName,
		}

		subCategory, err = qtx.CreateSubCategory(ctx, createSubCategoryParam)
		if err != nil {
			return uuid.Nil, err
		}
	}
	return subCategory.CategoryID, nil
}

func registerPostSubCategory(ctx context.Context, postId uuid.UUID, categoryId uuid.UUID, categoryNo string, qtx *Queries) error {
	createPostSubCategoryParam := CreatePostSubCategoryParams{
		PostID:     postId,
		CategoryID: categoryId,
		CategoryNo: categoryNo,
	}
	_, err := qtx.CreatePostSubCategory(ctx, createPostSubCategoryParam)
	if err != nil {
		return err
	}
	return nil
}

type CalcSummaryParam struct {
	Uid          string
	MainCategory string
	Expense      decimal.Decimal
	year         int32
	month        int32
	day          int32
}

func calcExpenseSummary(ctx context.Context, param CalcSummaryParam, qtx *Queries) error {
	getExpenseSummaryParam := GetExpenseSummaryParams{
		Uid:          param.Uid,
		Year:         param.year,
		Month:        param.month,
		MainCategory: param.MainCategory,
	}
	expenseSummary, err := qtx.GetExpenseSummary(ctx, getExpenseSummaryParam)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			createExpenseSummaryParam := CreateExpenseSummaryParams{
				Uid:          param.Uid,
				Year:         param.year,
				Month:        param.month,
				MainCategory: param.MainCategory,
				Amount:       param.Expense,
			}
			_, err := qtx.CreateExpenseSummary(ctx, createExpenseSummaryParam)
			if err != nil {
				return err
			}
		}
		return err
	}

	if expenseSummary.Amount.GreaterThan(MaxAmount) {
		return nil
	}

	updatedAmount := expenseSummary.Amount.Add(param.Expense)
	if updatedAmount.GreaterThan(MaxAmount) {
		updatedAmount = MaxAmount
	}

	updateExpenseSummaryParam := UpdateExpenseSummaryParams{
		Amount:       updatedAmount,
		Uid:          param.Uid,
		Year:         param.year,
		Month:        param.month,
		MainCategory: param.MainCategory,
	}

	_, err = qtx.UpdateExpenseSummary(ctx, updateExpenseSummaryParam)
	if err != nil {
		return err
	}
	return nil
}

func calcSubcategorySummary(ctx context.Context, param CalcSummaryParam, categoryId uuid.UUID, qtx *Queries) error {
	getSubcategorySummaryParam := GetSubCategorySummaryParams{
		Uid:        param.Uid,
		Year:       param.year,
		Month:      param.month,
		CategoryID: categoryId,
	}
	subcategorySummary, err := qtx.GetSubCategorySummary(ctx, getSubcategorySummaryParam)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			createSubcategorySummaryParam := CreateSubCategorySummaryParams{
				Uid:        param.Uid,
				Year:       param.year,
				Month:      param.month,
				CategoryID: categoryId,
				Amount:     param.Expense,
			}
			_, err := qtx.CreateSubCategorySummary(ctx, createSubcategorySummaryParam)
			if err != nil {
				return err
			}
		}
		return err
	}

	if subcategorySummary.Amount.GreaterThan(MaxAmount) {
		return nil
	}

	updatedAmount := subcategorySummary.Amount.Add(param.Expense)
	if updatedAmount.GreaterThan(MaxAmount) {
		updatedAmount = MaxAmount
	}

	updateSubcategorySummaryParam := UpdateSubCategorySummaryParams{
		Amount:     updatedAmount,
		Uid:        param.Uid,
		Year:       param.year,
		Month:      param.month,
		CategoryID: categoryId,
	}

	_, err = qtx.UpdateSubCategorySummary(ctx, updateSubcategorySummaryParam)
	if err != nil {
		return err
	}
	return nil
}

func calcDailyActivitySummary(ctx context.Context, param CalcSummaryParam, qtx *Queries) error {
	getDailyActivitySummaryParam := GetDailyActivitySummaryParams{
		Uid:   param.Uid,
		Year:  param.year,
		Month: param.month,
		Day:   param.day,
	}
	dailyActivitySummary, err := qtx.GetDailyActivitySummary(ctx, getDailyActivitySummaryParam)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			createDailyActivitySummaryParam := CreateDailyActivitySummaryParams{
				Uid:    param.Uid,
				Year:   param.year,
				Month:  param.month,
				Day:    param.day,
				Number: 1,
				Amount: param.Expense,
			}
			_, err := qtx.CreateDailyActivitySummary(ctx, createDailyActivitySummaryParam)
			if err != nil {
				return err
			}
		}
		return err
	}

	updateDailyActivitySummaryParam := UpdateDailyActivitySummaryParams{
		Number: dailyActivitySummary.Number + 1,
		Amount: dailyActivitySummary.Amount,
		Uid:    param.Uid,
		Year:   param.year,
		Month:  param.month,
		Day:    param.day,
	}

	if dailyActivitySummary.Amount.GreaterThan(MaxAmount) {
		qtx.UpdateDailyActivitySummary(ctx, updateDailyActivitySummaryParam)
	}

	updatedAmount := dailyActivitySummary.Amount.Add(param.Expense)
	if updatedAmount.GreaterThan(MaxAmount) {
		updatedAmount = MaxAmount
	}
	updateDailyActivitySummaryParam.Amount = updatedAmount

	_, err = qtx.UpdateDailyActivitySummary(ctx, updateDailyActivitySummaryParam)
	if err != nil {
		return err
	}
	return nil
}
