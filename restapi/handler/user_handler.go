package handler

import (
	"context"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kngnkg/tunetrail/restapi/model"
	"github.com/kngnkg/tunetrail/restapi/service"
)

type UserService interface {
	GetUserByUserName(ctx context.Context, userName string) (*model.User, error)
	UpdateUser(ctx context.Context, u *model.UserUpdateData) error
	DeleteUserByUserName(ctx context.Context, userName string) error
}

type UserHandler struct {
	Service UserService
}

// GET /user/:user_name
// ユーザー名からユーザーを取得する
func (uh *UserHandler) GetUserByUserName(c *gin.Context) {
	userName := c.Param("user_name")
	u, err := uh.Service.GetUserByUserName(c.Request.Context(), userName)
	if err != nil {
		c.Error(err)
		// ユーザーが存在しない場合
		if errors.Is(err, service.ErrUserNotFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"msg": UserNotFoundMessage})
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": ServerErrorMessage})
		return
	}
	c.JSON(http.StatusOK, u)
}

// PUT /user
// ユーザーを更新する
func (uh *UserHandler) UpdateUser(c *gin.Context) {
	var data *model.UserUpdateData
	if err := c.ShouldBindJSON(&data); err != nil {
		c.Error(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": BadRequestFieldMessage})
		return
	}
	err := uh.Service.UpdateUser(c.Request.Context(), data)
	if err != nil {
		c.Error(err)
		// ユーザーが存在しない場合
		if errors.Is(err, service.ErrUserNotFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"msg": UserNotFoundMessage})
			return
		}
		// ユーザー名が既に登録されている場合
		if errors.Is(err, service.ErrUserNameAlreadyExists) {
			c.AbortWithStatusJSON(http.StatusConflict, gin.H{"msg": UserNameAlreadyEntryMessage})
			return
		}
		// メールアドレスが既に登録されている場合
		if errors.Is(err, service.ErrEmailAlreadyExists) {
			c.AbortWithStatusJSON(http.StatusConflict, gin.H{"msg": EmailAlreadyEntryMessage})
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": ServerErrorMessage})
		return
	}
	c.Status(http.StatusNoContent)
}

// DELETE /user/:user_name
// ユーザーを削除する
func (uh *UserHandler) DeleteUserByUserName(c *gin.Context) {
	userName := c.Param("user_name")
	err := uh.Service.DeleteUserByUserName(c.Request.Context(), userName)
	if err != nil {
		c.Error(err)
		// ユーザーが存在しない場合
		if errors.Is(err, service.ErrUserNotFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"msg": UserNotFoundMessage})
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": ServerErrorMessage})
		return
	}
	c.Status(http.StatusNoContent)
}
