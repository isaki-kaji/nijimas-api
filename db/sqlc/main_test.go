package db_test

import (
	"os"
	"testing"

	db "github.com/isaki-kaji/nijimas-api/db/sqlc"
	"github.com/isaki-kaji/nijimas-api/testutil"
)

var testRepository db.Repository

func TestMain(m *testing.M) {
	testRepository = testutil.SetupDB()
	os.Exit(m.Run())
}
