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
}

type AuthHandler struct {
	Service AuthService
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

	c.JSON(http.StatusOK, tokens)
}
