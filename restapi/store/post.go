package store

import (
	"context"

	"github.com/kngnkg/tunetrail/restapi/model"
	"github.com/lib/pq"
)

func (r *Repository) GetPostsByUserIdsNext(ctx context.Context, db Queryer, userIds []model.UserID, pagenation *model.Pagenation) (*model.Timeline, error) {
	var posts []*model.Post

	baseQuery := `
		SELECT
			p.id,
			p.body,
			p.created_at,
			p.updated_at,
			u.id AS "user.id",
			u.user_name AS "user.user_name",
			u.name AS "user.name",
			u.icon_url AS "user.icon_url",
			u.bio AS "user.bio",
			u.created_at AS "user.created_at"
		FROM posts p
		INNER JOIN users u ON p.user_id = u.id
		WHERE user_id = ANY ($1)
	`

	limit := pagenation.Limit + 1 // 次のページがあるかどうかを判定するために1件多く取得する

	queryArgs := []interface{}{pq.Array(userIds), limit}

	var statement string

	if pagenation.NextCursor == "" {
		statement = baseQuery + `
			ORDER BY p.created_at DESC
			LIMIT $2;
		`
	} else {
		statement = baseQuery + `
			AND p.created_at <= (
				SELECT created_at
				FROM posts
				WHERE id = $3)
			ORDER BY p.created_at DESC
			LIMIT $2;
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

func (r *Repository) GetPostsByUserIdsPrevious(ctx context.Context, db Queryer, userId []model.UserID, previousCursor string) ([]*model.Post, error) {
	var posts []*model.Post

	statement := `
	SELECT id, user_id, body, created_at, updated_at
	FROM posts
	WHERE user_id = ANY ($1)
	AND created_at <= (
		SELECT created_at
		FROM posts
		WHERE id = $2)
	ORDER BY created_at ASC
	LIMIT 10;
	`

	err := db.SelectContext(ctx, &posts, statement, userId, previousCursor)
	if err != nil {
		return nil, err
	}

	return posts, nil
}

// TODO: idを返すようにする
func (r *Repository) AddPost(ctx context.Context, db Preparer, p *model.Post) (*model.Post, error) {
	p.CreatedAt = r.Clocker.Now()
	p.UpdatedAt = r.Clocker.Now()

	statement := `
		INSERT INTO posts (user_id, body, created_at, updated_at)
		VALUES ($1, $2, $3, $4)
		RETURNING id;
		`

	stmt, err := db.PreparexContext(ctx, statement)
	if err != nil {
		return nil, err
	}

	err = stmt.QueryRowxContext(ctx, p.User.Id, p.Body, p.CreatedAt, p.UpdatedAt).StructScan(p)
	if err != nil {
		return nil, err
	}

	return p, nil
}
