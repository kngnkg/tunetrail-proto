package handler

import (
	"context"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kngnkg/tunetrail/restapi/model"
	"github.com/kngnkg/tunetrail/restapi/service"
)

type AuthService interface {
	RegisterUser(ctx context.Context, data *model.UserRegistrationData) (*model.User, error)
	ConfirmEmail(ctx context.Context, userName, code string) error
	SignIn(ctx context.Context, data *model.UserSignInData) (*model.Tokens, error)
	GetSignedInUser(ctx context.Context, idToken string) (*model.User, error)
	RefreshToken(ctx context.Context, userIdentifier, refreshToken string) (string, error)
}

type AuthHandler struct {
	Service       AuthService
	AllowedDomain string
}

// POST /auth/register
// ユーザーを登録する
func (ah *AuthHandler) RegisterUser(c *gin.Context) {
	var data *model.UserRegistrationData
	if err := c.ShouldBindJSON(&data); err != nil {
		c.Error(err)
		errorResponse(c, http.StatusBadRequest, InvalidParameterCode)
		return
	}

	u, err := ah.Service.RegisterUser(c.Request.Context(), data)
	if err != nil {
		// ユーザー名が既に登録されている場合
		if errors.Is(err, service.ErrUserNameAlreadyExists) {
			errorResponse(c, http.StatusConflict, UserNameAlreadyEntryCode)
			return
		}
		// メールアドレスが既に登録されている場合
		if errors.Is(err, service.ErrEmailAlreadyExists) {
			errorResponse(c, http.StatusConflict, EmailAlreadyEntryCode)
			return
		}

		c.Error(err)
		errorResponse(c, http.StatusInternalServerError, ServerErrorCode)
		return
	}

	// ユーザーIDを返す
	c.JSON(http.StatusOK, gin.H{"id": u.Id})
}

// PUT /auth/confirm
// メールアドレスを確認する
func (ah *AuthHandler) ConfirmEmail(c *gin.Context) {
	var b struct {
		UserName string `json:"userName" binding:"required"`
		Code     string `json:"code" binding:"required"`
	}

	if err := c.ShouldBindJSON(&b); err != nil {
		c.Error(err)
		errorResponse(c, http.StatusBadRequest, InvalidParameterCode)
		return
	}

	if err := ah.Service.ConfirmEmail(c.Request.Context(), b.UserName, b.Code); err != nil {
		// ユーザー名が存在しない場合
		if errors.Is(err, service.ErrUserNotFound) {
			errorResponse(c, http.StatusNotFound, UserNotFoundCode)
			return
		}
		// メールアドレスの確認コードが不正な場合
		if errors.Is(err, service.ErrCodeMismatch) {
			errorResponse(c, http.StatusBadRequest, InvalidConfirmationCode)
			return
		}
		// メールアドレスの確認コードが期限切れの場合
		if errors.Is(err, service.ErrCodeExpired) {
			errorResponse(c, http.StatusBadRequest, ConfirmationCodeExpiredCode)
			return
		}
		// メールアドレスが既に確認済みの場合
		if errors.Is(err, service.ErrEmailAlreadyExists) {
			errorResponse(c, http.StatusConflict, EmailAlreadyConfirmedCode)
			return
		}

		c.Error(err)
		errorResponse(c, http.StatusInternalServerError, ServerErrorCode)
		return
	}

	c.Status(http.StatusNoContent)
}

// POST /auth/signin
// サインインする
func (ah *AuthHandler) SignIn(c *gin.Context) {
	var data *model.UserSignInData
	if err := c.ShouldBindJSON(&data); err != nil {
		c.Error(err)
		errorResponse(c, http.StatusBadRequest, InvalidParameterCode)
		return
	}
	if data.UserName == "" && data.Email == "" {
		errorResponse(c, http.StatusBadRequest, InvalidParameterCode)
		return
	}

	tokens, err := ah.Service.SignIn(c.Request.Context(), data)
	if err != nil {
		if errors.Is(err, service.ErrUserNotFound) {
			errorResponse(c, http.StatusNotFound, UserNotFoundCode)
			return
		}
		// メールアドレスまたはパスワードが一致しない場合
		if errors.Is(err, service.ErrNotAuthorized) {
			errorResponse(c, http.StatusBadRequest, NotAuthorizedCode)
			return
		}

		c.Error(err)
		errorResponse(c, http.StatusInternalServerError, ServerErrorCode)
		return
	}

	// HttpOnly Cookie にトークンをセット
	setCookie(c, IdTokenKey, tokens.Id)
	setCookie(c, AccessTokenKey, tokens.Access)
	setCookie(c, RefreshTokenKey, tokens.Refresh)

	// サインインしたユーザーの情報を取得する
	user, err := ah.Service.GetSignedInUser(c.Request.Context(), tokens.Id)
	if err != nil {
		c.Error(err)
		errorResponse(c, http.StatusInternalServerError, ServerErrorCode)
		return
	}

	c.JSON(http.StatusOK, user)
}

// POST /auth/refresh
// リフレッシュトークンを使ってアクセストークンを更新する
func (ah *AuthHandler) RefreshToken(c *gin.Context) {
	rt, err := c.Cookie(RefreshTokenKey)
	if err != nil {
		// エラーハンドリングを考える
		c.Error(err)
		errorResponse(c, http.StatusUnauthorized, NotAuthorizedCode)
		return
	}

	id := c.MustGet(UserIdKey).(string)

	accessToken, err := ah.Service.RefreshToken(c.Request.Context(), id, rt)
	if err != nil {
		// // リフレッシュトークンが不正な場合
		// if errors.Is(err, service.ErrInvalidRefreshToken) {
		// 	errorResponse(c, http.StatusBadRequest, InvalidRefreshTokenCode)
		// 	return
		// }
		// // リフレッシュトークンが期限切れの場合
		// if errors.Is(err, service.ErrRefreshTokenExpired) {
		// 	errorResponse(c, http.StatusBadRequest, RefreshTokenExpiredCode)
		// 	return
		// }

		c.Error(err)
		errorResponse(c, http.StatusInternalServerError, ServerErrorCode)
		return
	}

	setCookie(c, AccessTokenKey, accessToken)

	c.Status(http.StatusNoContent)
}

func setCookie(c *gin.Context, key, value string) {
	switch key {
	case IdTokenKey:
		c.SetCookie(
			key,
			value,
			60*60*24*7, // TODO: 有効期限を短くする
			"/",
			"."+c.Request.Host,
			true, // Secure
			true, // HttpOnly
		)
	case AccessTokenKey:
		c.SetCookie(
			key,
			value,
			60*60*24*7, // TODO: 有効期限を短くする
			"/",
			"."+c.Request.Host,
			true, // Secure
			true, // HttpOnly
		)
	case RefreshTokenKey:
		c.SetCookie(
			key,
			value,
			60*60*24*7, // TODO: 有効期限を短くする
			"/auth/refresh",
			"."+c.Request.Host,
			true, // Secure
			true, // HttpOnly
		)
	default:
		panic("invalid key")
	}
}
