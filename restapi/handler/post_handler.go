package handler

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kngnkg/tunetrail/restapi/model"
)

type PostService interface {
	AddPost(ctx context.Context, signedInUserId model.UserID, body string) (*model.Post, error)
	GetTimelines(ctx context.Context, signedInUserId model.UserID, pagenation *model.Pagenation) (*model.Timeline, error)
	GetPostsByUserId(ctx context.Context, userId model.UserID, pagenation *model.Pagenation) (*model.Timeline, error)
}

type PostHandler struct {
	Service PostService
}

// POST /posts
func (h *PostHandler) AddPost(c *gin.Context) {
	var b struct {
		Body string `json:"body" binding:"required"`
	}

	if err := c.ShouldBindJSON(&b); err != nil {
		c.Error(err)
		errorResponse(c, http.StatusBadRequest, InvalidParameterCode)
		return
	}

	signedInUserId := getSignedInUserId(c)

	p, err := h.Service.AddPost(c.Request.Context(), signedInUserId, b.Body)
	if err != nil {
		c.Error(err)
		errorResponse(c, http.StatusInternalServerError, ServerErrorCode)
		return
	}

	c.JSON(http.StatusCreated, p)
}

// GET /users/timelines
func (h *PostHandler) GetTimeline(c *gin.Context) {
	nc := c.DefaultQuery("next_cursor", "")
	pc := c.DefaultQuery("previous_cursor", "")
	lstr := c.DefaultQuery("limit", "10")

	l, err := strconv.Atoi(lstr)
	if err != nil {
		c.Error(err)
		errorResponse(c, http.StatusBadRequest, InvalidParameterCode)
		return
	}

	pagenation := &model.Pagenation{
		NextCursor:     nc,
		PreviousCursor: pc,
		Limit:          l,
	}

	signedInUserId := getSignedInUserId(c)

	timeline, err := h.Service.GetTimelines(c.Request.Context(), signedInUserId, pagenation)
	if err != nil {
		c.Error(err)
		errorResponse(c, http.StatusInternalServerError, ServerErrorCode)
		return
	}

	c.JSON(http.StatusOK, timeline)
}

func (h *PostHandler) GetPostsByUserId(c *gin.Context) {
	userId := model.UserID(c.Param("user_id"))

	nc := c.DefaultQuery("next_cursor", "")
	pc := c.DefaultQuery("previous_cursor", "")
	lstr := c.DefaultQuery("limit", "10")

	l, err := strconv.Atoi(lstr)
	if err != nil {
		c.Error(err)
		errorResponse(c, http.StatusBadRequest, InvalidParameterCode)
		return
	}

	pagenation := &model.Pagenation{
		NextCursor:     nc,
		PreviousCursor: pc,
		Limit:          l,
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
