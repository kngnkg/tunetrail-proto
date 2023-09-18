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

		if parentId != "" {
			err = ps.Repo.AddReplyRelation(ctx, ps.DB, id, parentId)

			if err != nil {
				return err
			}
		}

		registered, err := ps.Repo.GetPostById(ctx, ps.DB, id, signedInUserId)

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

	timeline, err := ps.Repo.GetPostsByUserIds(ctx, ps.DB, userIds, signedInUserId, pagenation)

	if err != nil {
		return nil, err
	}

	return timeline, nil
}

func (ps *PostService) GetPostsByUserId(ctx context.Context, userId model.UserID, signedInUserId model.UserID, pagenation *model.Pagenation) (*model.Timeline, error) {
	timeline, err := ps.Repo.GetPostsByUserId(ctx, ps.DB, userId, signedInUserId, pagenation)

	if err != nil {
		return nil, err
	}

	return timeline, nil
}

func (ps *PostService) GetPostById(ctx context.Context, postId string, signedInUserId model.UserID) (*model.Post, error) {
	p, err := ps.Repo.GetPostById(ctx, ps.DB, postId, signedInUserId)

	if err != nil {
		return nil, err
	}

	return p, nil
}

func (ps *PostService) GetReplies(ctx context.Context, postId string, pagenation *model.Pagenation) (*model.Timeline, error) {
	timeline, err := ps.Repo.GetReplies(ctx, ps.DB, postId, pagenation)

	if err != nil {
		return nil, err
	}

	for _, p := range timeline.Posts {
		if p.User.Id != "" {
			continue
		}

		// 削除されたポストの場合
		p.Body = "このポストは削除されました。"
	}

	return timeline, nil
}

func (ps *PostService) DeletePost(ctx context.Context, postId string) error {
	err := ps.Repo.WithTransaction(ctx, ps.DB, func(tx *sqlx.Tx) error {
		if err := ps.Repo.DeletePost(ctx, tx, postId); err != nil {
			return err
		}

		// ツリー構造が破壊されてしまうので、リプライ関係は削除しない

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}
