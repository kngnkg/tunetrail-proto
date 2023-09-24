package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/kngnkg/tunetrail/restapi/model"
	"github.com/kngnkg/tunetrail/restapi/store"
)

type HealthRepository interface {
	Ping(ctx context.Context, db store.Queryer) error
}

type HealthService struct {
	DB   store.DBConnection
	Repo HealthRepository
}

// HealthCheckは疎通を確認し、Health構造体を返す
func (hs *HealthService) HealthCheck(ctx context.Context) (*model.Health, error) {
	if err := hs.Repo.Ping(ctx, hs.DB); err != nil {
		if errors.Is(err, store.ErrCannotCommunicateWithDB) {
			h := &model.Health{
				Health:   model.StatusOrange,
				Database: model.StatusRed,
			}
			return h, err
		}
		h := &model.Health{
			Health:   model.StatusOrange,
			Database: model.StatusRed,
		}
		return h, fmt.Errorf("unexpected error: %w", err)
	}

	h := &model.Health{
		Health:   model.StatusGreen,
		Database: model.StatusGreen,
	}
	return h, nil
}
