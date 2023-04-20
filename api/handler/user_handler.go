package handler

import (
	"context"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kwtryo/tunetrail/api/model"
	"github.com/kwtryo/tunetrail/api/service"
)

type UserService interface {
	RegisterUser(
		ctx context.Context, userName, name, password, email, iconUrl, Bio string,
	) (*model.User, error)
	GetUserByUserName(ctx context.Context, userName string) (*model.User, error)
	DeleteUserByUserName(ctx context.Context, userName string) error
}

type UserHandler struct {
	Service UserService
}

// POST /user/register
// ユーザーを登録する
// TODO: ロガー関数を作成してログを出力する
func (uh *UserHandler) RegisterUser(c *gin.Context) {
	var req model.UserRegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": BadRequestMessage})
		return
	}
	u, err := uh.Service.RegisterUser(
		c.Request.Context(), req.UserName, req.Name, req.Password, req.Email, req.IconUrl, req.Bio,
	)
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

// GET /user/:user_name
// ユーザー名からユーザーを取得する
func (uh *UserHandler) GetUserByUserName(c *gin.Context) {
	userName := c.Param("user_name")
	u, err := uh.Service.GetUserByUserName(c.Request.Context(), userName)
	if err != nil {
		c.Error(err)
		// ユーザーが存在しない場合
		if errors.Is(err, service.ErrUserNotFound) {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": UserNotFoundMessage})
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": ServerErrorMessage})
		return
	}
	c.JSON(http.StatusOK, u)
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
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": UserNotFoundMessage})
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": ServerErrorMessage})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": SuccessMessage})
}
