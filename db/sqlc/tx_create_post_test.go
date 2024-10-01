package db_test

import (
	"context"
	"testing"
	"time"

	db "github.com/isaki-kaji/nijimas-api/db/sqlc"
	"github.com/isaki-kaji/nijimas-api/util"
)

func BenchmarkTxCreatePost(b *testing.B) {
	postId, err := util.GenerateUUID()
	if err != nil {
		b.Error(err)
	}

	createdAt, err := time.Parse("2006-01-02 15:04:05 -0700", "2023-09-30 14:45:00 +0000")
	if err != nil {
		b.Error(err)
	}

	testData := db.CreatePostTxParam{
		PostID:       postId,
		Uid:          "2TRw03nXoEg6CoecHZFsIoNLFjs2",
		MainCategory: "food",
		SubCategory1: "タコベル",
		SubCategory2: "メキシコ料理",
		PostText:     util.ToPointerOrNil("とても美味しかったです。"),
		Expense:      "1500",
		PublicTypeNo: "0",
		CreatedAt:    createdAt,
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := testRepository.CreatePostTx(context.Background(), testData)
		if err != nil {
			b.Error(err)
			break
		}

		b.StopTimer()
		resetDB(testData)
		b.StartTimer()
	}
}

func resetDB(param db.CreatePostTxParam) {
	ctx := context.Background()

	// deleteDailyActivitySummaryParams := db.DeleteDailyActivitySummaryParams{
	// 	Uid:   param.Uid,
	// 	Year:  2023,
	// 	Month: 9,
	// 	Day:   30,
	// }
	// err := testRepository.DeleteDailyActivitySummary(ctx, deleteDailyActivitySummaryParams)
	// if err != nil {
	// 	panic(err)
	// }

	// deleteExpenseSummaryParams := db.DeleteExpenseSummaryParams{
	// 	Uid:          param.Uid,
	// 	Year:         2023,
	// 	Month:        9,
	// 	MainCategory: param.MainCategory,
	// }
	// err = testRepository.DeleteExpenseSummary(ctx, deleteExpenseSummaryParams)
	// if err != nil {
	// 	panic(err)
	// }

	subCategory1, err := testRepository.GetSubCategoryByName(ctx, param.SubCategory1)
	if err != nil {
		panic(err)
	}
	subCategory2, err := testRepository.GetSubCategoryByName(ctx, param.SubCategory2)
	if err != nil {
		panic(err)
	}

	// deleteSubCategorySummaryParams1 := db.DeleteSubCategorySummaryParams{
	// 	Uid:        param.Uid,
	// 	Year:       2023,
	// 	Month:      9,
	// 	CategoryID: subCategory1.CategoryID,
	// }
	// err = testRepository.DeleteSubCategorySummary(ctx, deleteSubCategorySummaryParams1)
	// if err != nil {
	// 	panic(err)
	// }

	// deleteSubCategorySummaryParams2 := db.DeleteSubCategorySummaryParams{
	// 	Uid:        param.Uid,
	// 	Year:       2023,
	// 	Month:      9,
	// 	CategoryID: subCategory2.CategoryID,
	// }
	// err = testRepository.DeleteSubCategorySummary(ctx, deleteSubCategorySummaryParams2)
	// if err != nil {
	// 	panic(err)
	// }

	deletePostSubCategoryParams1 := db.DeletePostSubCategoryParams{
		PostID:     param.PostID,
		CategoryNo: "1",
	}
	err = testRepository.DeletePostSubCategory(ctx, deletePostSubCategoryParams1)
	if err != nil {
		panic(err)
	}

	deletePostSubCategoryParams2 := db.DeletePostSubCategoryParams{
		PostID:     param.PostID,
		CategoryNo: "2",
	}
	err = testRepository.DeletePostSubCategory(ctx, deletePostSubCategoryParams2)
	if err != nil {
		panic(err)
	}

	err = testRepository.DeleteSubCategory(ctx, subCategory1.CategoryID)
	if err != nil {
		panic(err)
	}
	err = testRepository.DeleteSubCategory(ctx, subCategory2.CategoryID)
	if err != nil {
		panic(err)
	}

	err = testRepository.DeletePost(ctx, param.PostID)
	if err != nil {
		panic(err)
	}
}

// 条件 VSCode以外のプロセスが動いていないこと

// 2024-09-30
// 直列処理
// BenchmarkTxCreatePost-8   	     788	   1336789 ns/op	   10135 B/op	     284 allocs/op
