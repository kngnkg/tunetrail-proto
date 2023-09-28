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
	GetPostById(ctx context.Context, db store.Queryer, postId string) (*model.Post, error)
	GetUserByUserId(ctx context.Context, db store.Queryer, id model.UserID) (*model.User, error)
	GetLikeInfoByPostId(ctx context.Context, db store.Queryer, signedInUserId model.UserID, postId string) (*model.LikeInfo, error)
}

type LikeService struct {
	DB   store.DBConnection
	Repo LikeRepository
}

func (ls *LikeService) AddLike(ctx context.Context, userId model.UserID, postId string) (*model.Post, error) {
	err := ls.Repo.WithTransaction(ctx, ls.DB, func(tx *sqlx.Tx) error {
		err := ls.Repo.AddLike(ctx, tx, userId, postId)

		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	likedPost, err := ls.Repo.GetPostById(ctx, ls.DB, postId)

	if err != nil {
		return nil, err
	}

	// ユーザー情報を紐付ける
	u, err := ls.Repo.GetUserByUserId(ctx, ls.DB, likedPost.User.Id)

	if err != nil {
		return nil, err
	}

	likedPost.User = u

	// いいね情報を紐付ける
	likeInfo, err := ls.Repo.GetLikeInfoByPostId(ctx, ls.DB, userId, likedPost.Id)

	if err != nil {
		return nil, err
	}

	likedPost.LikeInfo = *likeInfo

	return likedPost, nil
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
