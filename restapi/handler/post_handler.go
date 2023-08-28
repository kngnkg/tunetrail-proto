package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kngnkg/tunetrail/restapi/model"
)

type PostService interface {
	AddPost(ctx context.Context, signedInUserId model.UserID, body string) (*model.Post, error)
	// GetTimelines(ctx context.Context, signedInUserId model.UserID) (*model.Timeline, error)
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

// GET /:user_name/timelines
func (h *PostHandler) GetTimeline(c *gin.Context) {
	// signedInUserId := getSignedInUserId(c)

	// timeline, err := h.Service.GetTimelines(c.Request.Context(), signedInUserId)
	// if err != nil {
	// 	c.Error(err)
	// 	errorResponse(c, http.StatusInternalServerError, ServerErrorCode)
	// 	return
	// }

	timeline := &model.Timeline{
		Posts: []*model.Post{
			{
				Id:        "1",
				UserId:    "1",
				Body:      "body",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			{
				Id:        "2",
				UserId:    "2",
				Body:      "body2",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			{
				Id:        "3",
				UserId:    "3",
				Body:      "body3",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
		},
		Pagenation: &model.Pagenation{
			NextCursor:     "next",
			PreviousCursor: "previous",
			Limit:          10,
		},
	}

	c.JSON(http.StatusOK, timeline)
}
