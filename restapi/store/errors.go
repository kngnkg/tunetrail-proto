package store

import "errors"

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
