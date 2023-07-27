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
