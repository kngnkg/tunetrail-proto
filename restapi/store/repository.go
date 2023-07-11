package store

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/kngnkg/tunetrail/restapi/clock"
	"github.com/kngnkg/tunetrail/restapi/config"
	_ "github.com/lib/pq"
)

// Postgresのエラーコード
const (
	// 重複エラーコード
	ErrCodePostgresDuplicate = "23505"
)

// storeパッケージで用いるエラー
var (
	// DBとの疎通が取れない
	ErrCannotCommunicateWithDB = errors.New("store: cannot communicate with db")
	// トランザクション開始に失敗
	ErrBeginTxFailed = errors.New("store: begin tx failed")
	// ロールバックに失敗
	ErrRollbackFailed = errors.New("store: rollback failed")
	// コミットに失敗
	ErrCommitFailed = errors.New("store: commit failed")
	// ユーザーが見つからない
	ErrUserNotFound = errors.New("store: user not found")
	// ユーザー名が既に存在する
	ErrUserNameAlreadyExists = errors.New("store: user name already exists")
	// メールアドレスが既に存在する
	ErrEmailAlreadyExists = errors.New("store: email already exists")
)

type Repository struct {
	Clocker clock.Clocker
}

// NewはconfigからDBオブジェクトを返す
func New(cfg *config.Config) (*sqlx.DB, func(), error) {
	driver := "postgres"

	sslMode := "require"
	if cfg.Env == "dev" {
		sslMode = "disable" // 開発環境の場合はSSL通信を無効にする
	}

	db, err := sql.Open(driver, fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName, sslMode,
	))
	if err != nil {
		log.Printf("store: failed to open db: %v", err)
		return nil, nil, err
	}

	// sql.Openは接続確認が行われないため、ここで確認する。
	if err := db.Ping(); err != nil {
		log.Printf("store: failed to ping db: %v", err)
		return nil, func() { _ = db.Close() }, err
	}
	xdb := sqlx.NewDb(db, driver)
	return xdb, func() { _ = db.Close() }, nil
}

// トランザクションを実行するためのインターフェース
type Beginner interface {
	BeginTxx(ctx context.Context, opts *sql.TxOptions) (*sqlx.Tx, error)
}

type Preparer interface {
	PreparexContext(ctx context.Context, query string) (*sqlx.Stmt, error)
}

// Queryerはsqlx.DBとsqlx.Txのインターフェースを統一するためのインターフェース
type Queryer interface {
	Preparer
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
	NamedExecContext(ctx context.Context, query string, arg interface{}) (sql.Result, error)
	QueryxContext(ctx context.Context, query string, args ...any) (*sqlx.Rows, error)
	QueryRowxContext(ctx context.Context, query string, args ...any) *sqlx.Row
	GetContext(ctx context.Context, dest interface{}, query string, args ...any) error
	SelectContext(ctx context.Context, dest interface{}, query string, args ...any) error
}

// インターフェースが期待通りに宣言されているか確認
var (
	_ Beginner = (*sqlx.DB)(nil)
	_ Preparer = (*sqlx.DB)(nil)
	_ Queryer  = (*sqlx.DB)(nil)
	_ Queryer  = (*sqlx.Tx)(nil)
)

// WithTransactionはトランザクションを実行する
func (r *Repository) WithTransaction(ctx context.Context, db Beginner, f func(tx *sqlx.Tx) error) error {
	tx, err := db.BeginTxx(ctx, nil)
	if err != nil {
		return fmt.Errorf("%w: %w", ErrBeginTxFailed, err)
	}

	defer func() {
		// トランザクション内でpanicが発生した場合はRollbackを実行する
		if p := recover(); p != nil {
			if err := tx.Rollback(); err != nil {
				panic(fmt.Errorf("%w: %w", ErrRollbackFailed, err))
			}
			panic(p)
		}
	}()

	// トランザクション内で実行する処理
	err = f(tx)
	// トランザクション内でエラーが発生した場合はRollbackを実行する
	if err != nil {
		if err := tx.Rollback(); err != nil {
			return fmt.Errorf("%w: %w", ErrRollbackFailed, err)
		}
		return err
	}

	// トランザクションをコミットする
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("%w: %w", ErrCommitFailed, err)
	}

	return nil
}
