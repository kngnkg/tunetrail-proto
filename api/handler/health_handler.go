package handler

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kwtryo/tunetrail/api/model"
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
		log.Printf("ERROR: %v", err)
		c.JSON(http.StatusInternalServerError, h)
		return
	}
	c.JSON(http.StatusOK, h)
}
