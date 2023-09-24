package service

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/kngnkg/tunetrail/restapi/model"
	"github.com/kngnkg/tunetrail/restapi/store"
)

type LikeRepository interface {
	Transactioner
	AddLike(ctx context.Context, db store.Execer, userId model.UserID, postId string) error
	DeleteLike(ctx context.Context, db store.Execer, userId model.UserID, postId string) error
	GetPostById(ctx context.Context, db store.Queryer, postId string, signedInUserId model.UserID) (*model.Post, error)
}

type LikeService struct {
	DB   store.DBConnection
	Repo LikeRepository
}

func (ls *LikeService) AddLike(ctx context.Context, userId model.UserID, postId string) (*model.Post, error) {
	var p *model.Post

	err := ls.Repo.WithTransaction(ctx, ls.DB, func(tx *sqlx.Tx) error {
		err := ls.Repo.AddLike(ctx, tx, userId, postId)

		if err != nil {
			return err
		}

		likedPost, err := ls.Repo.GetPostById(ctx, ls.DB, postId, userId)

		if err != nil {
			return err
		}

		p = likedPost
		return nil
	})

	if err != nil {
		return nil, err
	}

	return p, nil
}

func (ls *LikeService) DeleteLike(ctx context.Context, userId model.UserID, postId string) error {
	err := ls.Repo.WithTransaction(ctx, ls.DB, func(tx *sqlx.Tx) error {
		err := ls.Repo.DeleteLike(ctx, tx, userId, postId)

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
