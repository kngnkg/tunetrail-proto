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
			p.parent_id,
			u.id AS "user.id",
			u.user_name AS "user.user_name",
			u.name AS "user.name",
			u.icon_url AS "user.icon_url",
			u.bio AS "user.bio",
			u.created_at AS "user.created_at",
			u.updated_at AS "user.updated_at",
			p.body,
			p.created_at,
			p.updated_at
		FROM posts p
		INNER JOIN users u ON p.user_id = u.id
		WHERE user_id = ANY ($1)
		AND p.parent_id IS NULL;
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

	tl := handlePagenation(posts, pagenation)

	return tl, nil
}

func (r *Repository) GetPostById(ctx context.Context, db Queryer, postId string) (*model.Post, error) {
	p := &model.Post{}

	statement := `
		SELECT
			p.id,
			COALESCE(CAST(p.parent_id AS text), '') AS "parent_id", -- NULLの場合にGoのstring型にバインドできないため文字列に変換する
			u.id AS "user.id",
			u.user_name AS "user.user_name",
			u.name AS "user.name",
			u.icon_url AS "user.icon_url",
			u.bio AS "user.bio",
			u.created_at AS "user.created_at",
			u.updated_at AS "user.updated_at",
			p.body,
			p.created_at,
			p.updated_at
		FROM posts p
		INNER JOIN users u ON p.user_id = u.id
		WHERE p.id = $1;
	`

	if err := db.GetContext(ctx, p, statement, postId); err != nil {
		return nil, err
	}

	return p, nil
}

// 昇順で取得する
func (r *Repository) GetReplies(ctx context.Context, db Queryer, parentPostId string, pagenation *model.Pagenation) (*model.Timeline, error) {
	var posts []*model.Post

	baseQuery := `
	WITH RECURSIVE post_tree AS (
		-- ベースケース
		SELECT
			r1.post_id AS "id",
			r1.parent_id AS "parent_id",
			r1.created_at AS "reply_created_at",
			p.body,
			p.created_at,
			p.updated_at,
			u.id AS "user_id",
			u.user_name AS "user_user_name",
			u.name AS "user_name",
			u.icon_url AS "user_icon_url",
			u.bio AS "user_bio",
			u.created_at AS "user_created_at",
			u.updated_at AS "user_updated_at"
		FROM
			replies r1 -- ツリー構造をたどるためにrepliesテーブルを左テーブルとして結合する
		LEFT OUTER JOIN posts p ON r1.post_id = p.id -- 削除された投稿の場合はNULLになる
		LEFT OUTER JOIN users u ON p.user_id = u.id
	WHERE
		r1.parent_id = $1 -- 最初のリプライを取得する

	UNION ALL

	-- 再帰的ケース
	SELECT
		r2.post_id AS "id",
		r2.parent_id AS "parent_id",
		r2.created_at AS "reply_created_at",
		p.body,
		p.created_at,
		p.updated_at,
		u.id AS "user_id",
		u.user_name AS "user_user_name",
		u.name AS "user_name",
		u.icon_url AS "user_icon_url",
		u.bio AS "user_bio",
		u.created_at AS "user_created_at",
		u.updated_at AS "user_updated_at"
	FROM
		replies r2
		LEFT OUTER JOIN posts p ON r2.post_id = p.id
		LEFT OUTER JOIN users u ON p.user_id = u.id
		INNER JOIN post_tree pt ON r2.parent_id = pt.id
	)
	SELECT
		COALESCE(CAST(id AS text), '') AS "id", -- NULLの場合にGoのstring型にバインドできないため文字列に変換する
		COALESCE(CAST(parent_id AS text), '') AS "parent_id",
		COALESCE(body, '') AS "body",
		COALESCE(created_at, now()) AS "created_at",
		COALESCE(updated_at, now()) AS "updated_at",
		COALESCE(CAST(user_id AS text), '') AS "user.id",
		COALESCE(user_user_name, '') AS "user.user_name",
		COALESCE(user_name, '') AS "user.name",
		COALESCE(user_icon_url, '') AS "user.icon_url",
		COALESCE(user_bio, '') AS "user.bio",
		COALESCE(user_created_at, now()) AS "user.created_at",
		COALESCE(user_updated_at, now()) AS "user.updated_at"
	FROM
		post_tree
	`

	limit := pagenation.Limit + 1 // 次のページがあるかどうかを判定するために1件多く取得する

	queryArgs := []interface{}{parentPostId, limit}

	var statement string

	if pagenation.NextCursor == "" {
		statement = baseQuery + `
			ORDER BY reply_created_at ASC -- リプライの作成日時の昇順で取得する
			LIMIT $2;
		`
	} else {
		statement = baseQuery + `
			AND reply_created_at <= (
				SELECT created_at
				FROM replies
				WHERE post_id = $3)
			ORDER BY reply_created_at ASC
			LIMIT $2;
		`

		queryArgs = append(queryArgs, pagenation.NextCursor)
	}

	if err := db.SelectContext(ctx, &posts, statement, queryArgs...); err != nil {
		return nil, err
	}

	tl := handlePagenation(posts, pagenation)

	return tl, nil
}

func (r *Repository) AddPost(ctx context.Context, db Queryer, p *model.Post) (string, error) {
	p.CreatedAt = r.Clocker.Now()
	p.UpdatedAt = r.Clocker.Now()

	var statement string
	queryArgs := []interface{}{p.User.Id, p.Body, p.CreatedAt, p.UpdatedAt}

	if p.ParentId == "" {
		statement = `
			INSERT INTO posts (user_id, body, created_at, updated_at)
			VALUES($1, $2, $3, $4)
			RETURNING id;
		`
	} else {
		// リプライの場合
		statement = `
			INSERT INTO posts (user_id, parent_id, body, created_at, updated_at)
			VALUES($1, $5, $2, $3, $4)
			RETURNING id;
		`

		queryArgs = append(queryArgs, p.ParentId)
	}

	var id string
	err := db.QueryRowxContext(ctx, statement, queryArgs...).Scan(&id)

	if err != nil {
		return "", err
	}

	return id, nil
}

func handlePagenation(posts []*model.Post, pagenation *model.Pagenation) *model.Timeline {
	limit := pagenation.Limit + 1 // 次のページがあるかどうかを判定するために1件多く取得する

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

	return tl
}
