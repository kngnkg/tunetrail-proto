package service

import "errors"

var (
	// ユーザー名が既に存在する
	ErrUserNameAlreadyExists = errors.New("service: user name already exists")
	// メールアドレスが既に存在する
	ErrEmailAlreadyExists = errors.New("service: email already exists")
	// ユーザーが存在しない
	ErrUserNotFound = errors.New("service: user not found")
	// 検証コードが一致しない
	ErrCodeMismatch = errors.New("service: code mismatch")
	// 検証コードが期限切れ
	ErrCodeExpired = errors.New("service: code expired")
)
