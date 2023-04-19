package handler

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/kwtryo/tunetrail/api/model"
	"github.com/kwtryo/tunetrail/api/service"
)

type UserService interface {
	RegisterUser(
		ctx context.Context, userName, name, password, email, iconUrl, Bio string,
	) (*model.User, error)
}

type UserHandler struct {
	Service UserService
}

// POST /user/register
// ユーザーを登録する
// TODO: ロガー関数を作成してログを出力する
func (uh *UserHandler) RegisterUser(c *gin.Context) {
	// バリデーションの初期化
	initValidation()

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
		if err == service.ErrEmailAlreadyExists {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": EmailAlreadyEntryMessage})
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": ServerErrorMessage})
		return
	}
	// ユーザーIDを返す
	c.JSON(http.StatusOK, gin.H{"id": u.Id})
}

// TODO: ユーザーのバリデーションを行う関数を別ファイルに切り出す
func initValidation() {
	// カスタムバリデーションルールを登録
	validate := binding.Validator.Engine().(*validator.Validate)
	err := validate.RegisterValidation("password", model.PasswordValidationFunction)
	if err != nil {
		fmt.Printf("Failed to register custom validation: %v\n", err)
		return
	}
}
