package store

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/kwtryo/tunetrail/api/clock"
	"github.com/kwtryo/tunetrail/api/config"
	_ "github.com/lib/pq"
)

// storeパッケージで用いるエラー
var (
	ErrNotFound     = errors.New("not found")
	ErrAlreadyEntry = errors.New("duplicate entry")
)

type Repository struct {
	Clocker clock.Clocker
}

// NewはconfigからDBオブジェクトを返す
func New(cfg *config.Config) (*sqlx.DB, func(), error) {
	driver := "postgres"
	db, err := sql.Open(driver, fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser,
		cfg.DBPassword, cfg.DBName,
	))
	if err != nil {
		return nil, nil, err
	}

	// sql.Openは接続確認が行われないため、ここで確認する。
	if err := db.Ping(); err != nil {
		return nil, func() { _ = db.Close() }, err
	}
	xdb := sqlx.NewDb(db, driver)
	return xdb, func() { _ = db.Close() }, nil
}

type Beginner interface {
	BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error)
}

type Preparer interface {
	PreparexContext(ctx context.Context, query string) (*sqlx.Stmt, error)
}

type Execer interface {
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
	NamedExecContext(ctx context.Context, query string, arg interface{}) (sql.Result, error)
}

type Queryer interface {
	Preparer
	QueryxContext(ctx context.Context, query string, args ...any) (*sqlx.Rows, error)
	QueryRowxContext(ctx context.Context, query string, args ...any) *sqlx.Row
	GetContext(ctx context.Context, dest interface{}, query string, args ...any) error
	SelectContext(ctx context.Context, dest interface{}, query string, args ...any) error
}

// インターフェースが期待通りに宣言されているか確認
var (
	_ Beginner = (*sqlx.DB)(nil)
	_ Preparer = (*sqlx.DB)(nil)
	_ Execer   = (*sqlx.DB)(nil)
	_ Execer   = (*sqlx.Tx)(nil)
	_ Queryer  = (*sqlx.DB)(nil)
)
