package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/kngnkg/tunetrail/restapi/model"
	"github.com/kngnkg/tunetrail/restapi/store"
)

type FollowRepository interface {
	Transactioner
	GetUserByUserId(ctx context.Context, db store.Queryer, id model.UserID) (*model.User, error)
	AddFollow(ctx context.Context, db store.Execer, userId, follweeUserId model.UserID) error
	DeleteFollow(ctx context.Context, db store.Execer, userId, follweeUserId model.UserID) error
}

type FollowService struct {
	DB   store.DBConnection
	Repo FollowRepository
}

// FollowUserはユーザーをフォローする
func (fs *FollowService) FollowUser(ctx context.Context, userId, follweeUserId model.UserID) (*model.User, error) {
	var u *model.User

	err := fs.Repo.WithTransaction(ctx, fs.DB, func(tx *sqlx.Tx) error {
		if err := fs.Repo.AddFollow(ctx, tx, userId, follweeUserId); err != nil {
			if errors.Is(err, store.ErrUserNotFound) {
				return fmt.Errorf("%w: userId=%v: %w", ErrUserNotFound, userId, err)
			}
			return err
		}

		got, err := fs.Repo.GetUserByUserId(ctx, tx, follweeUserId)

		if err != nil {
			return err
		}

		u = got

		return nil
	})

	if err != nil {
		return nil, err
	}

	return u, nil
}

// UnfollowUserはユーザーのフォローを解除する
func (fs *FollowService) UnfollowUser(ctx context.Context, userId, follweeUserId model.UserID) error {
	err := fs.Repo.WithTransaction(ctx, fs.DB, func(tx *sqlx.Tx) error {
		if err := fs.Repo.DeleteFollow(ctx, tx, userId, follweeUserId); err != nil {
			if errors.Is(err, store.ErrUserNotFound) {
				return fmt.Errorf("%w: userId=%v: %w", ErrUserNotFound, userId, err)
			}
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}
