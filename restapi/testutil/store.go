package testutil

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"reflect"
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/kngnkg/tunetrail/restapi/model"
	"github.com/stretchr/testify/assert"
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

// Password、Email以外のフィールドが一致することを確認する。
func AssertUser(t *testing.T, expected, actial *model.User) {
	t.Helper()

	if expected == nil {
		assert.Nil(t, actial)
		return
	}

	ev := reflect.ValueOf(expected).Elem()
	av := reflect.ValueOf(actial).Elem()
	for i := 0; i < ev.NumField(); i++ {
		field := ev.Type().Field(i)
		if field.Name == "Password" || field.Name == "Email" {
			continue
		}
		assert.Equal(t, ev.Field(i).Interface(), av.Field(i).Interface())
	}
}

func DeletePostAll(t *testing.T, ctx context.Context, tx *sqlx.Tx) {
	t.Helper()

	if _, err := tx.ExecContext(ctx, "DELETE FROM posts"); err != nil {
		t.Fatal(err)
	}
}

func DeleteFollowAll(t *testing.T, ctx context.Context, tx *sqlx.Tx) {
	t.Helper()

	if _, err := tx.ExecContext(ctx, "DELETE FROM follows"); err != nil {
		t.Fatal(err)
	}
}
