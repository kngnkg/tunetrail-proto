package store

import (
	"context"

	"github.com/kngnkg/tunetrail/restapi/model"
	"github.com/lib/pq"
)

func (r *Repository) GetPostsByUserIdsNext(ctx context.Context, db Queryer, userIds []model.UserID, pagenation *model.Pagenation) (*model.Timeline, error) {
	var posts []*model.Post

	baseQuery := `
		SELECT id, user_id, body, created_at, updated_at
		FROM posts
		WHERE user_id = ANY ($1)
	`

	limit := pagenation.Limit + 1 // 次のページがあるかどうかを判定するために1件多く取得する

	queryArgs := []interface{}{pq.Array(userIds), limit}

	var statement string

	if pagenation.NextCursor == "" {
		statement = baseQuery + `
			ORDER BY created_at DESC
			LIMIT $2;
		`
	} else {
		statement = baseQuery + `
			AND created_at <= (
				SELECT created_at
				FROM posts
				WHERE id = $2)
			ORDER BY created_at DESC
			LIMIT $3;
		`

		queryArgs = append(queryArgs, pagenation.NextCursor)
	}

	if err := db.SelectContext(ctx, &posts, statement, queryArgs...); err != nil {
		return nil, err
	}

	if len(posts) == limit {
		// 次のページがある場合は、次のページのためにカーソルをセットする
		pagenation.NextCursor = posts[limit-1].Id
		posts = posts[:limit-1]
	} else {
		pagenation.NextCursor = ""
	}

	tl := &model.Timeline{
		Posts:      posts,
		Pagenation: pagenation,
	}

	return tl, nil
}

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
