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
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": BadRequestFieldMessage})
		return
	}
	u, err := ah.Service.RegisterUser(c.Request.Context(), data)
	if err != nil {
		c.Error(err)
		// ユーザー名が既に登録されている場合
		if errors.Is(err, service.ErrUserNameAlreadyExists) {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": UserNameAlreadyEntryMessage})
			return
		}
		// メールアドレスが既に登録されている場合
		if errors.Is(err, service.ErrEmailAlreadyExists) {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": EmailAlreadyEntryMessage})
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": ServerErrorMessage})
		return
	}
	// ユーザーIDを返す
	c.JSON(http.StatusOK, gin.H{"id": u.Id})
}
