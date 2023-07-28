package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Code             int    `json:"code"`
	DeveloperMessage string `json:"developerMessage"`
	UserMessage      string `json:"userMessage"`
}

func errorResponse(c *gin.Context, statusCode, errorCode int) {
	msg, ok := ErrorMessages[errorCode]
	if !ok {
		statusCode = http.StatusInternalServerError
		msg = ErrorMessage{
			DeveloperMessage: "An unknown error occurred",
			UserMessage:      "An unknown error occurred",
		}
	}

	body := ErrorResponse{
		Code:             errorCode,
		DeveloperMessage: msg.DeveloperMessage,
		UserMessage:      msg.UserMessage,
	}
	c.AbortWithStatusJSON(statusCode, body)
}

// エラーコード

// 汎用的なエラーコード
const (
	BadRequestCode = 4000 + iota
	InvalidParameterCode
)

// 認証関連のエラーコード
const (
	InvalidConfirmationCode = 4101 + iota
	ConfirmationCodeExpiredCode
	EmailAlreadyConfirmedCode
)

// ユーザー関連のエラーコード
const (
	UserNotFoundCode = 4201 + iota
	UserNameAlreadyEntryCode
	EmailAlreadyEntryCode
)

// その他サーバー内部のエラーコード
const (
	ServerErrorCode = 5000
)

// エラーメッセージ

type ErrorMessage struct {
	DeveloperMessage string
	UserMessage      string
}

var ErrorMessages = map[int]ErrorMessage{
	// 汎用的なエラー
	BadRequestCode: {
		DeveloperMessage: "Bad request",
		UserMessage:      "不明なエラーが発生しました。",
	},
	InvalidParameterCode: {
		DeveloperMessage: "Invalid parameter",
		UserMessage:      "不明なエラーが発生しました。",
	},
	// 認証関連
	InvalidConfirmationCode: {
		DeveloperMessage: "Invalid confirmation code",
		UserMessage:      "確認コードが一致しません。",
	},
	ConfirmationCodeExpiredCode: {
		DeveloperMessage: "Confirmation code expired",
		UserMessage:      "確認コードが期限切れです。",
	},
	EmailAlreadyConfirmedCode: {
		DeveloperMessage: "Email already confirmed",
		UserMessage:      "既に確認済みのメールアドレスです。",
	},
	// ユーザー関連
	UserNotFoundCode: {
		DeveloperMessage: "User not found",
		UserMessage:      "ユーザーが存在しません。",
	},
	UserNameAlreadyEntryCode: {
		DeveloperMessage: "UserName already entry",
		UserMessage:      "ユーザー名が既に存在します。",
	},
	EmailAlreadyEntryCode: {
		DeveloperMessage: "Email already entry",
		UserMessage:      "登録できないメールアドレスです。",
	},
	// サーバー内部のエラー
	ServerErrorCode: {
		DeveloperMessage: "Unknown server error",
		UserMessage:      "不明なエラーが発生しました。",
	},
}
