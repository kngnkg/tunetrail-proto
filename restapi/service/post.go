package service

import (
	"context"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/kngnkg/tunetrail/restapi/model"
	"github.com/kngnkg/tunetrail/restapi/store"
)

type PostService struct {
	DB   store.DBConnection
	Repo PostRepository
}

func (ps *PostService) AddPost(ctx context.Context, signedInUserId model.UserID, body string) (*model.Post, error) {
	var p *model.Post
	err := ps.Repo.WithTransaction(ctx, ps.DB, func(tx *sqlx.Tx) error {

		added, err := ps.Repo.AddPost(ctx, ps.DB, &model.Post{
			UserId: signedInUserId,
			Body:   body,
		})

		if err != nil {
			return err
		}

		p = added
		return nil
	})

	if err != nil {
		return nil, err
	}

	return p, nil
}

func (ps *PostService) GetTimelines(ctx context.Context, signedInUserId model.UserID, pagenation *model.Pagenation) (*model.Timeline, error) {
	// ユーザーがフォローしているユーザーのIDを取得する
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

	for _, u := range userIds {
		log.Printf("userId: %v", u)
	}

	// フィードを取得する
	timeline, err := ps.Repo.GetPostsByUserIdsNext(ctx, ps.DB, userIds, pagenation)

	if err != nil {
		return nil, err
	}

	return timeline, nil
}
