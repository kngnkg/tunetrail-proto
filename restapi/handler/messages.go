package handler

// レスポンスメッセージ
const (
	SuccessMessage                 = "成功しました。"
	BadRequestMessage              = "不正なリクエストです。"
	BadRequestFieldMessage         = "不正なフィールドがあります。"
	ServerErrorMessage             = "サーバー内部でエラーが発生しました。"
	UserNotFoundMessage            = "ユーザーが存在しません。"
	UserNameAlreadyEntryMessage    = "ユーザー名が既に登録されています。"
	EmailAlreadyEntryMessage       = "メールアドレスが既に登録されています。"
	InvalidConfirmationCodeMessage = "メールアドレスの確認コードが不正です。"
	ConfirmationCodeExpiredMessage = "メールアドレスの確認コードが期限切れです。"
)
