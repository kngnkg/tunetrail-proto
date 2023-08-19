package store

import (
	"context"

	"github.com/kngnkg/tunetrail/restapi/model"
)

func (r *Repository) AddPost(ctx context.Context, db Preparer, p *model.Post) (*model.Post, error) {
	p.CreatedAt = r.Clocker.Now()
	p.UpdatedAt = r.Clocker.Now()

	statement := `INSERT INTO posts (user_id, body, created_at, updated_at)
				VALUES ($1, $2, $3, $4)
				RETURNING id, user_id, body, created_at, updated_at;`

	stmt, err := db.PreparexContext(ctx, statement)
	if err != nil {
		return nil, err
	}

	err = stmt.QueryRowxContext(ctx, p.UserId, p.Body, p.CreatedAt, p.UpdatedAt).StructScan(p)
	if err != nil {
		return nil, err
	}

	return p, nil
}
