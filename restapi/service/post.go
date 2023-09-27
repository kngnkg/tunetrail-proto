package service

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/kngnkg/tunetrail/restapi/model"
	"github.com/kngnkg/tunetrail/restapi/store"
	"golang.org/x/sync/errgroup"
)

type PostRepository interface {
	Transactioner
	GetFolloweesByUserId(ctx context.Context, db store.Queryer, signedInUserId model.UserID) ([]*model.User, error)
	AddPost(ctx context.Context, db store.Queryer, p *model.Post) (string, error)
	AddReplyRelation(ctx context.Context, db store.Execer, postId, parentId string) error
	DeletePost(ctx context.Context, db store.Execer, postId string) error
	GetPostById(ctx context.Context, db store.Queryer, postId string, signedInUserId model.UserID) (*model.Post, error)
	GetPostsByUserId(ctx context.Context, db store.Queryer, userId model.UserID, signedInUserId model.UserID, pagination *model.Pagination) (*model.Timeline, error)
	GetPostsByUserIds(ctx context.Context, db store.Queryer, userIds []model.UserID, signedInUserId model.UserID, pagination *model.Pagination) (*model.Timeline, error)
	GetLikedPostsByUserId(ctx context.Context, db store.Queryer, userId model.UserID, signedInUserId model.UserID, pagination *model.Pagination) (*model.Timeline, error)
	GetReplies(ctx context.Context, db store.Queryer, parentPostId string, pagination *model.Pagination) (*model.Timeline, error)
	GetUserByUserId(ctx context.Context, db store.Queryer, id model.UserID) (*model.User, error)
}

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

func (ps *PostService) GetTimelines(ctx context.Context, signedInUserId model.UserID, pagination *model.Pagination) (*model.Timeline, error) {
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

	timeline, err := ps.Repo.GetPostsByUserIds(ctx, ps.DB, userIds, signedInUserId, pagination)

	if err != nil {
		return nil, err
	}

	posts, err := ps.fillUserToPosts(ctx, timeline.Posts)

	if err != nil {
		return nil, err
	}

	timeline.Posts = posts

	return timeline, nil
}

func (ps *PostService) GetPostsByUserId(ctx context.Context, userId model.UserID, signedInUserId model.UserID, pagination *model.Pagination) (*model.Timeline, error) {
	timeline, err := ps.Repo.GetPostsByUserId(ctx, ps.DB, userId, signedInUserId, pagination)

	if err != nil {
		return nil, err
	}

	posts, err := ps.fillUserToPosts(ctx, timeline.Posts)

	if err != nil {
		return nil, err
	}

	timeline.Posts = posts

	return timeline, nil
}

func (ps *PostService) GetLikedPostsByUserId(ctx context.Context, userId model.UserID, signedInUserId model.UserID, pagination *model.Pagination) (*model.Timeline, error) {
	timeline, err := ps.Repo.GetLikedPostsByUserId(ctx, ps.DB, userId, signedInUserId, pagination)

	if err != nil {
		return nil, err
	}

	posts, err := ps.fillUserToPosts(ctx, timeline.Posts)

	if err != nil {
		return nil, err
	}

	timeline.Posts = posts

	return timeline, nil
}

func (ps *PostService) GetPostById(ctx context.Context, postId string, signedInUserId model.UserID) (*model.Post, error) {
	p, err := ps.Repo.GetPostById(ctx, ps.DB, postId, signedInUserId)

	if err != nil {
		return nil, err
	}

	// ユーザー情報を紐付ける
	u, err := ps.Repo.GetUserByUserId(ctx, ps.DB, p.User.Id)

	if err != nil {
		return nil, err
	}

	p.User = u

	return p, nil
}

func (ps *PostService) GetReplies(ctx context.Context, postId string, pagination *model.Pagination) (*model.Timeline, error) {
	timeline, err := ps.Repo.GetReplies(ctx, ps.DB, postId, pagination)

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

// ユーザー情報を紐付ける
func (ps *PostService) fillUserToPosts(ctx context.Context, posts []*model.Post) ([]*model.Post, error) {
	eg, ctx := errgroup.WithContext(ctx)

	for _, p := range posts {
		post := p

		eg.Go(func() error {
			u, err := ps.Repo.GetUserByUserId(ctx, ps.DB, post.User.Id)

			if err != nil {
				return err
			}

			post.User = u
			return nil
		})
	}

	if err := eg.Wait(); err != nil {
		return nil, err
	}

	return posts, nil
}
