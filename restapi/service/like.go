package service

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/kngnkg/tunetrail/restapi/model"
	"github.com/kngnkg/tunetrail/restapi/store"
)

type LikeService struct {
	DB   store.DBConnection
	Repo LikeRepository
}

func (ls *LikeService) AddLike(ctx context.Context, userId model.UserID, postId string) error {
	err := ls.Repo.WithTransaction(ctx, ls.DB, func(tx *sqlx.Tx) error {
		err := ls.Repo.AddLike(ctx, tx, userId, postId)

		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}
