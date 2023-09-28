package handler

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kngnkg/tunetrail/restapi/model"
)

type PostService interface {
	AddPost(ctx context.Context, signedInUserId model.UserID, ParentId, body string) (*model.Post, error)
	GetPostById(ctx context.Context, postId string, signedInUserId model.UserID) (*model.Post, error)
	GetTimelines(ctx context.Context, signedInUserId model.UserID, pagination *model.Pagination) ([]*model.Post, *model.Pagination, error)
	GetPostsByUserId(ctx context.Context, userId model.UserID, signedInUserId model.UserID, pagination *model.Pagination) ([]*model.Post, *model.Pagination, error)
	GetLikedPostsByUserId(ctx context.Context, userId model.UserID, signedInUserId model.UserID, pagination *model.Pagination) ([]*model.Post, *model.Pagination, error)
	GetReplies(ctx context.Context, postId string, signedInUserId model.UserID, pagination *model.Pagination) ([]*model.Post, *model.Pagination, error)
	DeletePost(ctx context.Context, postId string) error
}

type postsResp struct {
	Posts      []*model.Post     `json:"posts"`
	Pagination *model.Pagination `json:"pagination"`
}

type PostHandler struct {
	Service PostService
}

func (h *PostHandler) AddPost(c *gin.Context) {
	var b struct {
		ParentId string `json:"parentId"`
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

func (h *PostHandler) GetTimeline(c *gin.Context) {
	signedInUserId := getSignedInUserId(c)

	reqPagination, err := getPaginationFromQuery(c)
	if err != nil {
		c.Error(err)
		errorResponse(c, http.StatusBadRequest, InvalidParameterCode)
		return
	}

	posts, pagination, err := h.Service.GetTimelines(c.Request.Context(), signedInUserId, reqPagination)
	if err != nil {
		c.Error(err)
		errorResponse(c, http.StatusInternalServerError, ServerErrorCode)
		return
	}

	c.JSON(http.StatusOK, &postsResp{
		Posts:      posts,
		Pagination: pagination,
	})
}

func (h *PostHandler) GetPostsByUserId(c *gin.Context) {
	userId := getUserIdFromPath(c)
	signedInUserId := getSignedInUserId(c)

	reqPagination, err := getPaginationFromQuery(c)
	if err != nil {
		c.Error(err)
		errorResponse(c, http.StatusBadRequest, InvalidParameterCode)
		return
	}

	posts, pagination, err := h.Service.GetPostsByUserId(c.Request.Context(), userId, signedInUserId, reqPagination)
	if err != nil {
		c.Error(err)
		errorResponse(c, http.StatusInternalServerError, ServerErrorCode)
		return
	}

	c.JSON(http.StatusOK, &postsResp{
		Posts:      posts,
		Pagination: pagination,
	})
}

func (h *PostHandler) GetLikedPostsByUserId(c *gin.Context) {
	userId := getUserIdFromPath(c)
	signedInUserId := getSignedInUserId(c)

	reqPagination, err := getPaginationFromQuery(c)
	if err != nil {
		c.Error(err)
		errorResponse(c, http.StatusBadRequest, InvalidParameterCode)
		return
	}

	posts, pagination, err := h.Service.GetLikedPostsByUserId(c.Request.Context(), userId, signedInUserId, reqPagination)
	if err != nil {
		c.Error(err)
		errorResponse(c, http.StatusInternalServerError, ServerErrorCode)
		return
	}

	c.JSON(http.StatusOK, &postsResp{
		Posts:      posts,
		Pagination: pagination,
	})
}

func (h *PostHandler) GetReplies(c *gin.Context) {
	postId := getPostIdFromPath(c)
	signInUserId := getSignedInUserId(c)

	reqPagination, err := getPaginationFromQuery(c)
	if err != nil {
		c.Error(err)
		errorResponse(c, http.StatusBadRequest, InvalidParameterCode)
		return
	}

	posts, pagination, err := h.Service.GetReplies(c.Request.Context(), postId, signInUserId, reqPagination)
	if err != nil {
		c.Error(err)
		errorResponse(c, http.StatusInternalServerError, ServerErrorCode)
		return
	}

	c.JSON(http.StatusOK, &postsResp{
		Posts:      posts,
		Pagination: pagination,
	})
}

func (h *PostHandler) GetPostById(c *gin.Context) {
	postId := getPostIdFromPath(c)
	signInUserId := getSignedInUserId(c)

	p, err := h.Service.GetPostById(c.Request.Context(), postId, signInUserId)
	if err != nil {
		c.Error(err)
		errorResponse(c, http.StatusInternalServerError, ServerErrorCode)
		return
	}

	c.JSON(http.StatusOK, p)
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
