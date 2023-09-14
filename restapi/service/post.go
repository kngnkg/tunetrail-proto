package service

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/kngnkg/tunetrail/restapi/model"
	"github.com/kngnkg/tunetrail/restapi/store"
)

type PostService struct {
	DB   store.DBConnection
	Repo PostRepository
}

func (ps *PostService) AddPost(ctx context.Context, signedInUserId model.UserID, parentId string, body string) (*model.Post, error) {
	var p *model.Post
	err := ps.Repo.WithTransaction(ctx, ps.DB, func(tx *sqlx.Tx) error {
		id, err := ps.Repo.AddPost(ctx, ps.DB, &model.Post{
			ParentId: parentId,
			User:     &model.User{Id: signedInUserId},
			Body:     body,
		})

		if err != nil {
			return err
		}

		registered, err := ps.Repo.GetPostById(ctx, ps.DB, id)

		if err != nil {
			return err
		}

		p = registered
		return nil
	})

	if err != nil {
		return nil, err
	}

	return p, nil
}

func (ps *PostService) GetTimelines(ctx context.Context, signedInUserId model.UserID, pagenation *model.Pagenation) (*model.Timeline, error) {
	// ユーザーがフォローしているユーザーの情報を取得する
	users, err := ps.Repo.GetFolloweesByUserId(ctx, ps.DB, signedInUserId)

	if err != nil {
		return nil, err
	}

	// ログインユーザーを含めたユーザーIDのスライスを作成する
	userIds := make([]model.UserID, len(users)+1)
	userIds[0] = signedInUserId
	for i, u := range users {
		userIds[i+1] = u.Id
	}

	// フィードを取得する
	timeline, err := ps.Repo.GetPostsByUserIdsNext(ctx, ps.DB, userIds, pagenation)

	if err != nil {
		return nil, err
	}

	return timeline, nil
}

func (ps *PostService) GetPostsByUserId(ctx context.Context, userId model.UserID, pagenation *model.Pagenation) (*model.Timeline, error) {
	// ユーザーIDのスライスを作成する
	userIds := make([]model.UserID, 1)
	userIds[0] = userId

	// フィードを取得する
	timeline, err := ps.Repo.GetPostsByUserIdsNext(ctx, ps.DB, userIds, pagenation)

	if err != nil {
		return nil, err
	}

	return timeline, nil
}
