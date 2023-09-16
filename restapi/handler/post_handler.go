package handler

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kngnkg/tunetrail/restapi/model"
)

type PostService interface {
	AddPost(ctx context.Context, signedInUserId model.UserID, ParentId, body string) (*model.Post, error)
	GetTimelines(ctx context.Context, signedInUserId model.UserID, pagenation *model.Pagenation) (*model.Timeline, error)
	GetPostsByUserId(ctx context.Context, userId model.UserID, pagenation *model.Pagenation) (*model.Timeline, error)
	GetPostById(ctx context.Context, postId string) (*model.Post, error)
	GetReplies(ctx context.Context, postId string, pagenation *model.Pagenation) (*model.Timeline, error)
	DeletePost(ctx context.Context, postId string) error
}

type PostHandler struct {
	Service PostService
}

// POST /posts
func (h *PostHandler) AddPost(c *gin.Context) {
	var b struct {
		ParentId string `json:"parent_id"`
		Body     string `json:"body" binding:"required"`
	}

	if err := c.ShouldBindJSON(&b); err != nil {
		c.Error(err)
		errorResponse(c, http.StatusBadRequest, InvalidParameterCode)
		return
	}

	signedInUserId := getSignedInUserId(c)

	p, err := h.Service.AddPost(c.Request.Context(), signedInUserId, b.ParentId, b.Body)
	if err != nil {
		c.Error(err)
		errorResponse(c, http.StatusInternalServerError, ServerErrorCode)
		return
	}

	c.JSON(http.StatusCreated, p)
}

// GET /users/timelines
func (h *PostHandler) GetTimeline(c *gin.Context) {
	signedInUserId := getSignedInUserId(c)

	pagenation, err := getPagenationFromQuery(c)
	if err != nil {
		c.Error(err)
		errorResponse(c, http.StatusBadRequest, InvalidParameterCode)
		return
	}

	timeline, err := h.Service.GetTimelines(c.Request.Context(), signedInUserId, pagenation)
	if err != nil {
		c.Error(err)
		errorResponse(c, http.StatusInternalServerError, ServerErrorCode)
		return
	}

	c.JSON(http.StatusOK, timeline)
}

func (h *PostHandler) GetPostsByUserId(c *gin.Context) {
	userId := getUserIdFromPath(c)

	pagenation, err := getPagenationFromQuery(c)
	if err != nil {
		c.Error(err)
		errorResponse(c, http.StatusBadRequest, InvalidParameterCode)
		return
	}

	// TODO: Timeline構造体の名前を変える
	timeline, err := h.Service.GetPostsByUserId(c.Request.Context(), userId, pagenation)
	if err != nil {
		c.Error(err)
		errorResponse(c, http.StatusInternalServerError, ServerErrorCode)
		return
	}

	c.JSON(http.StatusOK, timeline)
}

func (h *PostHandler) GetPostById(c *gin.Context) {
	postId := getPostIdFromPath(c)

	p, err := h.Service.GetPostById(c.Request.Context(), postId)
	if err != nil {
		c.Error(err)
		errorResponse(c, http.StatusInternalServerError, ServerErrorCode)
		return
	}

	c.JSON(http.StatusOK, p)
}

func (h *PostHandler) GetReplies(c *gin.Context) {
	postId := getPostIdFromPath(c)

	pagenation, err := getPagenationFromQuery(c)
	if err != nil {
		c.Error(err)
		errorResponse(c, http.StatusBadRequest, InvalidParameterCode)
		return
	}

	timeline, err := h.Service.GetReplies(c.Request.Context(), postId, pagenation)
	if err != nil {
		c.Error(err)
		errorResponse(c, http.StatusInternalServerError, ServerErrorCode)
		return
	}

	c.JSON(http.StatusOK, timeline)
}

func (h *PostHandler) DeletePost(c *gin.Context) {
	postId := getPostIdFromPath(c)

	if err := h.Service.DeletePost(c.Request.Context(), postId); err != nil {
		c.Error(err)
		errorResponse(c, http.StatusInternalServerError, ServerErrorCode)
		return
	}

	c.Status(http.StatusNoContent)
}
