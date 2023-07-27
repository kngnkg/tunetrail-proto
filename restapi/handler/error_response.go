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
	ServerErrorCode = 1000 + iota
	BadRequestCode
	BadRequestFieldCode
)

// ユーザー・認証関連のエラーコード
const (
	UserNotFoundCode = 2000 + iota
	UserNameAlreadyEntryCode
	EmailAlreadyEntryCode
	InvalidConfirmationCode
	ConfirmationCodeExpiredCode
)

// エラーメッセージ

type ErrorMessage struct {
	DeveloperMessage string
	UserMessage      string
}

var ErrorMessages = map[int]ErrorMessage{
	ServerErrorCode: {
		DeveloperMessage: "サーバー内部でエラーが発生しました。",
		UserMessage:      "不明なエラーが発生しました。",
	},
	BadRequestCode: {
		DeveloperMessage: "不正なリクエストです。",
		UserMessage:      "不正なリクエストです。",
	},
	BadRequestFieldCode: {
		DeveloperMessage: "不正なフィールドがあります。",
		UserMessage:      "不明なエラーが発生しました。",
	},
	UserNotFoundCode: {
		DeveloperMessage: "ユーザーが存在しません。",
		UserMessage:      "ユーザーが存在しません。",
	},
	UserNameAlreadyEntryCode: {
		DeveloperMessage: "ユーザー名が既に登録されています。",
		UserMessage:      "ユーザー名が既に存在します。",
	},
	EmailAlreadyEntryCode: {
		DeveloperMessage: "メールアドレスが既に登録されています。",
		UserMessage:      "登録できないメールアドレスです。",
	},
	InvalidConfirmationCode: {
		DeveloperMessage: "確認コードが一致しません。",
		UserMessage:      "確認コードが一致しません。",
	},
	ConfirmationCodeExpiredCode: {
		DeveloperMessage: "確認コードが期限切れです。",
		UserMessage:      "確認コードが期限切れです。",
	},
}

// ユーザー用レスポンスメッセージ
const (
	SuccessMessage                 = "成功しました。"
	ServerErrorMessage             = "サーバー内部でエラーが発生しました。"
	BadRequestMessage              = "不正なリクエストです。"
	BadRequestFieldMessage         = "不正なフィールドがあります。"
	UserNotFoundMessage            = "ユーザーが存在しません。"
	UserNameAlreadyEntryMessage    = "ユーザー名が既に登録されています。"
	EmailAlreadyEntryMessage       = "メールアドレスが既に登録されています。"
	InvalidConfirmationCodeMessage = "メールアドレスの確認コードが不正です。"
	ConfirmationCodeExpiredMessage = "メールアドレスの確認コードが期限切れです。"
)
