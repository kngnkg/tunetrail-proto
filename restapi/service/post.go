package service

import (
	"context"

	"github.com/kngnkg/tunetrail/restapi/model"
	"github.com/kngnkg/tunetrail/restapi/store"
)

type PostService struct {
	DB   store.DBConnection
	Repo PostRepository
}

func (ps *PostService) AddPost(ctx context.Context, signedInUserId model.UserID, body string) (*model.Post, error) {
	p := &model.Post{
		UserId: signedInUserId,
		Body:   body,
	}

	added, err := ps.Repo.AddPost(ctx, ps.DB, p)
	if err != nil {
		return nil, err
	}

	return added, nil
}
