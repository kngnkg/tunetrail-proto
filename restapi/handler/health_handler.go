package handler

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kngnkg/tunetrail/restapi/model"
)

type HealthService interface {
	HealthCheck(ctx context.Context) (*model.Health, error)
}

type HealthHandler struct {
	Service HealthService
}

// GET /health
// ヘルスチェック
func (hh *HealthHandler) HealthCheck(c *gin.Context) {
	h, err := hh.Service.HealthCheck(c.Request.Context())
	if err != nil {
		// エラーをコンテキストにセットする
		c.Error(err)
		// エラーが発生した場合は500エラーを返す
		c.JSON(http.StatusInternalServerError, h)
		return
	}
	c.JSON(http.StatusOK, h)
}
