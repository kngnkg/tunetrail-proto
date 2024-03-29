package handler

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kngnkg/tunetrail/restapi/model"
)

type LikeService interface {
	AddLike(ctx context.Context, userId model.UserID, postId string) (*model.Post, error)
	DeleteLike(ctx context.Context, userId model.UserID, postId string) error
}

type LikeHandler struct {
	Service LikeService
}

func (h *LikeHandler) AddLike(c *gin.Context) {
	signedInUserId := getSignedInUserId(c)
	postId := getPostIdFromPath(c)

	likedPost, err := h.Service.AddLike(c.Request.Context(), signedInUserId, postId)
	if err != nil {
		c.Error(err)
		errorResponse(c, http.StatusBadRequest, ServerErrorCode)
		return
	}

	c.JSON(http.StatusCreated, likedPost)
}

func (h *LikeHandler) DeleteLike(c *gin.Context) {
	signedInUserId := getSignedInUserId(c)
	postId := getPostIdFromPath(c)

	err := h.Service.DeleteLike(c.Request.Context(), signedInUserId, postId)
	if err != nil {
		c.Error(err)
		errorResponse(c, http.StatusBadRequest, ServerErrorCode)
		return
	}

	c.Status(http.StatusNoContent)
}
