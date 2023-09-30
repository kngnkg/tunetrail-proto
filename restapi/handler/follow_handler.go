package handler

import (
	"context"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kngnkg/tunetrail/restapi/model"
	"github.com/kngnkg/tunetrail/restapi/service"
)

type FollowService interface {
	FollowUser(ctx context.Context, userId, follweeUserId model.UserID) (*model.User, error)
	UnfollowUser(ctx context.Context, userId, follweeUserId model.UserID) error
}

type FollowHandler struct {
	Service FollowService
}

func (fh *FollowHandler) FollowUser(c *gin.Context) {
	userId := getUserIdFromPath(c)

	var b struct {
		FollweeUserId string `json:"followee_user_id"`
	}

	if err := c.ShouldBindJSON(&b); err != nil {
		c.Error(err)
		errorResponse(c, http.StatusBadRequest, InvalidParameterCode)
		return
	}

	u, err := fh.Service.FollowUser(c.Request.Context(), userId, model.UserID(b.FollweeUserId))
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

func (fh *FollowHandler) UnfollowUser(c *gin.Context) {
	userId := getUserIdFromPath(c)
	followeeId := getFolloweeIdFromPath(c)

	if err := fh.Service.UnfollowUser(c.Request.Context(), userId, followeeId); err != nil {
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
