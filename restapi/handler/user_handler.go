package handler

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kngnkg/tunetrail/restapi/model"
	"github.com/kngnkg/tunetrail/restapi/service"
)

type UserService interface {
	GetSignedInUser(ctx context.Context, userId model.UserID) (*model.User, error)
	GetUserByUserName(ctx context.Context, userName string, signedInUserId model.UserID) (*model.User, error)
	UpdateUser(ctx context.Context, u *model.UserUpdateData) error
	DeleteUserByUserName(ctx context.Context, userName string) error
	FollowUser(ctx context.Context, userId, follweeUserId model.UserID) (*model.User, error)
	UnfollowUser(ctx context.Context, userId, follweeUserId model.UserID) error
	GetFollowees(ctx context.Context, userId model.UserID) ([]*model.User, error)
	GetFollowers(ctx context.Context, userId model.UserID) ([]*model.User, error)
}

type UserHandler struct {
	Service UserService
}

// ログインユーザー情報取得
func (uh *UserHandler) GetMe(c *gin.Context) {
	signedInUserId := getSignedInUserId(c)

	u, err := uh.Service.GetSignedInUser(c.Request.Context(), signedInUserId)
	if err != nil {
		// ユーザーが存在しない場合
		if errors.Is(err, service.ErrUserNotFound) {
			errorResponse(c, http.StatusNotFound, UserNotFoundCode)
			return
		}

		c.Error(err)
		errorResponse(c, http.StatusInternalServerError, ServerErrorCode)
		return
	}

	c.JSON(http.StatusOK, u)
}

// GET /users/by/username/:user_name
// ユーザー名からユーザーを取得する
func (uh *UserHandler) GetUserByUserName(c *gin.Context) {
	userName := c.Param("user_name")

	signedInUserId := getSignedInUserId(c)

	u, err := uh.Service.GetUserByUserName(c.Request.Context(), userName, signedInUserId)
	if err != nil {
		// ユーザーが存在しない場合
		if errors.Is(err, service.ErrUserNotFound) {
			errorResponse(c, http.StatusNotFound, UserNotFoundCode)
			return
		}

		c.Error(err)
		errorResponse(c, http.StatusInternalServerError, ServerErrorCode)
		return
	}

	log.Println(u)

	c.JSON(http.StatusOK, u)
}

// ユーザーを更新する
func (uh *UserHandler) UpdateUser(c *gin.Context) {
	var data *model.UserUpdateData
	if err := c.ShouldBindJSON(&data); err != nil {
		c.Error(err)
		errorResponse(c, http.StatusBadRequest, InvalidParameterCode)
		return
	}
	err := uh.Service.UpdateUser(c.Request.Context(), data)
	if err != nil {
		// ユーザーが存在しない場合
		if errors.Is(err, service.ErrUserNotFound) {
			errorResponse(c, http.StatusNotFound, UserNotFoundCode)
			return
		}
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

	c.Status(http.StatusNoContent)
}

// DELETE /users/:user_id
// ユーザーを削除する
func (uh *UserHandler) DeleteUserByUserName(c *gin.Context) {
	userName := c.Param("user_name")
	err := uh.Service.DeleteUserByUserName(c.Request.Context(), userName)
	if err != nil {
		// ユーザーが存在しない場合
		if errors.Is(err, service.ErrUserNotFound) {
			errorResponse(c, http.StatusNotFound, UserNotFoundCode)
			return
		}

		c.Error(err)
		errorResponse(c, http.StatusInternalServerError, ServerErrorCode)
		return
	}

	c.Status(http.StatusNoContent)
}

// ユーザーをフォローする
func (uh *UserHandler) FollowUser(c *gin.Context) {
	userId := getUserIdFromPath(c)

	var b struct {
		FollweeUserId string `json:"followee_user_id"`
	}

	if err := c.ShouldBindJSON(&b); err != nil {
		c.Error(err)
		errorResponse(c, http.StatusBadRequest, InvalidParameterCode)
		return
	}

	u, err := uh.Service.FollowUser(c.Request.Context(), userId, model.UserID(b.FollweeUserId))
	if err != nil {
		// ユーザーが存在しない場合
		if errors.Is(err, service.ErrUserNotFound) {
			errorResponse(c, http.StatusNotFound, UserNotFoundCode)
			return
		}

		c.Error(err)
		errorResponse(c, http.StatusInternalServerError, ServerErrorCode)
		return
	}

	c.JSON(http.StatusCreated, u)
}

// ユーザーのフォローを解除する
func (uh *UserHandler) UnfollowUser(c *gin.Context) {
	userId := getUserIdFromPath(c)
	followeeId := getFolloweeIdFromPath(c)

	if err := uh.Service.UnfollowUser(c.Request.Context(), userId, followeeId); err != nil {
		// ユーザーが存在しない場合
		if errors.Is(err, service.ErrUserNotFound) {
			errorResponse(c, http.StatusNotFound, UserNotFoundCode)
			return
		}

		c.Error(err)
		errorResponse(c, http.StatusInternalServerError, ServerErrorCode)
		return
	}

	c.Status(http.StatusNoContent)
}

func (uh *UserHandler) GetFollowees(c *gin.Context) {
	userId := getUserIdFromPath(c)

	users, err := uh.Service.GetFollowees(c.Request.Context(), userId)
	if err != nil {
		// ユーザーが存在しない場合
		if errors.Is(err, service.ErrUserNotFound) {
			errorResponse(c, http.StatusNotFound, UserNotFoundCode)
			return
		}

		c.Error(err)
		errorResponse(c, http.StatusInternalServerError, ServerErrorCode)
		return
	}

	var response struct {
		Users []*model.User `json:"users"`
	}

	response.Users = users

	c.JSON(http.StatusOK, response)
}

func (uh *UserHandler) GetFollowers(c *gin.Context) {
	userId := getUserIdFromPath(c)

	users, err := uh.Service.GetFollowers(c.Request.Context(), userId)
	if err != nil {
		// ユーザーが存在しない場合
		if errors.Is(err, service.ErrUserNotFound) {
			errorResponse(c, http.StatusNotFound, UserNotFoundCode)
			return
		}

		c.Error(err)
		errorResponse(c, http.StatusInternalServerError, ServerErrorCode)
		return
	}

	var response struct {
		Users []*model.User `json:"users"`
	}

	response.Users = users

	c.JSON(http.StatusOK, response)
}
