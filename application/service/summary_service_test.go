package service_test

import (
	"context"
	"os"
	"testing"

	"github.com/isaki-kaji/nijimas-api/application/service"
	"github.com/isaki-kaji/nijimas-api/testutil"
)

var sSer service.SummaryService

func TestMain(m *testing.M) {
	testRepository := testutil.SetupDB()
	sSer = service.NewSummaryService(testRepository)
	os.Exit(m.Run())
}

func BenchmarkGetMonthlySummary(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := sSer.GetMonthlySummary(context.Background(), "2TRw03nXoEg6CoecHZFsIoNLFjs2", 2024, 10)
		if err != nil {
			b.Log("Error:", err)
			b.Error(err)
			break
		}
	}
}

// 条件 VSCode以外のプロセスが動いていないこと

// 2024-10-03
// 直列処理
// BenchmarkGetMonthlySummary-8   	    1454	    700784 ns/op	   25290 B/op	     741 allocs/op

// 2024-10-03
// 並行処理
// BenchmarkGetMonthlySummary-8   	    3151	    349857 ns/op	   26022 B/op	     750 allocs/op
