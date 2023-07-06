package testutil

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"testing"

	"github.com/jmoiron/sqlx"
)

// OpenDbForTestはテスト用のDBオブジェクトを返す。
func OpenDbForTest(t *testing.T, ctx context.Context) *sqlx.DB {
	t.Helper()

	cfg := CreateConfigForTest(t)
	driver := "postgres"
	db, err := sql.Open(driver, fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
	))
	if err != nil {
		log.Fatal(err)
	}
	t.Cleanup(
		func() { _ = db.Close() },
	)
	return sqlx.NewDb(db, driver)
}

// DeleteUserAllはテスト用のDBから全てのユーザーを削除する。
func DeleteUserAll(t *testing.T, ctx context.Context, tx *sqlx.Tx) {
	t.Helper()

	if _, err := tx.ExecContext(ctx, "DELETE FROM users"); err != nil {
		t.Fatal(err)
	}
}
