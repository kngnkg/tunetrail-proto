package handler

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kngnkg/tunetrail/restapi/model"
)

type PostService interface {
	AddPost(ctx context.Context, signedInUserId model.UserID, body string) (*model.Post, error)
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
