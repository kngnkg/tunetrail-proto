package store

import (
	"context"

	"github.com/kngnkg/tunetrail/restapi/model"
	"github.com/lib/pq"
)

const selectBasePostQuery = `
	SELECT
		p.id,
		COALESCE(CAST(r.parent_id AS text), '') AS "parent_id", -- NULLの場合にGoのstring型にバインドできないため文字列に変換する
		p.user_id AS "user.id",
		p.body,
		p.created_at,
		p.updated_at
	FROM
		posts p
		LEFT OUTER JOIN reply_relations r ON p.id = r.post_id
	`

func handlePagination(posts []*model.Post, pagination *model.Pagination) ([]*model.Post, *model.Pagination) {
	limit := pagination.Limit + 1 // 次のページがあるかどうかを判定するために1件多く取得する

	if len(posts) == limit {
		// 次のページがある場合は、次のページのためにカーソルをセットする
		pagination.NextCursor = posts[limit-1].Id
		posts = posts[:limit-1]

		return posts, pagination
	}

	pagination.NextCursor = ""
	return posts, pagination
}

func (r *Repository) GetPostsByUserIds(ctx context.Context, db Queryer, userIds []model.UserID, pagination *model.Pagination) ([]*model.Post, *model.Pagination, error) {
	var ps []*model.Post

	limit := pagination.Limit + 1 // 次のページがあるかどうかを判定するために1件多く取得する

	queryArgs := []interface{}{pq.Array(userIds), limit}

	statement := selectBasePostQuery + `
		WHERE p.user_id = ANY ($1) AND r.parent_id IS NULL
	`

	if pagination.NextCursor != "" {
		statement = statement + `
			AND p.created_at <= (SELECT created_at FROM posts WHERE id = $3)
		`

		queryArgs = append(queryArgs, pagination.NextCursor)
	}

	statement = statement + `
		ORDER BY p.created_at DESC
		LIMIT $2;
	`

	if err := db.SelectContext(ctx, &ps, statement, queryArgs...); err != nil {
		return nil, nil, err
	}

	posts, pg := handlePagination(ps, pagination)

	return posts, pg, nil
}

func (r *Repository) GetPostsByUserId(ctx context.Context, db Queryer, userId model.UserID, pagination *model.Pagination) ([]*model.Post, *model.Pagination, error) {
	var ps []*model.Post

	limit := pagination.Limit + 1 // 次のページがあるかどうかを判定するために1件多く取得する

	queryArgs := []interface{}{userId, limit}

	statement := selectBasePostQuery + `
		WHERE p.user_id = $1
	`

	if pagination.NextCursor != "" {
		statement = statement + `
			AND p.created_at <= (SELECT created_at FROM posts WHERE id = $3)
		`

		queryArgs = append(queryArgs, pagination.NextCursor)
	}

	statement = statement + `
		ORDER BY p.created_at DESC
		LIMIT $2;
	`

	if err := db.SelectContext(ctx, &ps, statement, queryArgs...); err != nil {
		return nil, nil, err
	}

	posts, pg := handlePagination(ps, pagination)

	return posts, pg, nil
}

func (r *Repository) GetLikedPostsByUserId(ctx context.Context, db Queryer, userId model.UserID, pagination *model.Pagination) ([]*model.Post, *model.Pagination, error) {
	var ps []*model.Post

	limit := pagination.Limit + 1 // 次のページがあるかどうかを判定するために1件多く取得する

	queryArgs := []interface{}{userId, limit}

	statement := selectBasePostQuery + `
		LEFT OUTER JOIN likes l ON p.id = l.post_id
		WHERE l.user_id = $1
	`

	if pagination.NextCursor != "" {
		statement = statement + `
			AND p.created_at <= (SELECT created_at FROM posts WHERE id = $3)
		`

		queryArgs = append(queryArgs, pagination.NextCursor)
	}

	statement = statement + `
		ORDER BY p.created_at DESC
		LIMIT $2;
	`

	if err := db.SelectContext(ctx, &ps, statement, queryArgs...); err != nil {
		return nil, nil, err
	}

	posts, pg := handlePagination(ps, pagination)

	return posts, pg, nil
}

func (r *Repository) GetPostById(ctx context.Context, db Queryer, postId string) (*model.Post, error) {
	p := &model.Post{}

	statement := selectBasePostQuery + `WHERE p.id = $1` + `;`

	if err := db.GetContext(ctx, p, statement, postId); err != nil {
		return nil, err
	}

	return p, nil
}

// 昇順で取得する
func (r *Repository) GetReplies(ctx context.Context, db Queryer, parentPostId string, pagination *model.Pagination) ([]*model.Post, *model.Pagination, error) {
	var ps []*model.Post

	baseQuery := `
	WITH RECURSIVE post_tree AS (
		-- ベースケース
		SELECT
			r1.post_id AS "id",
			r1.parent_id AS "parent_id",
			r1.created_at AS "reply_created_at",
			p.user_id AS "user_id",
			p.body,
			p.created_at,
			p.updated_at
		FROM
			reply_relations r1 -- ツリー構造をたどるためにreply_relationsテーブルを左テーブルとして結合する
			LEFT OUTER JOIN posts p ON r1.post_id = p.id -- 削除された投稿の場合はNULLになる
	WHERE
		r1.parent_id = $1 -- 最初のリプライを取得する

	UNION ALL

	-- 再帰的ケース
	SELECT
		r2.post_id AS "id",
		r2.parent_id AS "parent_id",
		r2.created_at AS "reply_created_at",
		p.user_id AS "user_id",
		p.body,
		p.created_at,
		p.updated_at
	FROM
		reply_relations r2
		LEFT OUTER JOIN posts p ON r2.post_id = p.id
		INNER JOIN post_tree pt ON r2.parent_id = pt.id
	)
	SELECT
		COALESCE(CAST(id AS text), '') AS "id", -- NULLの場合にGoのstring型にバインドできないため文字列に変換する
		COALESCE(CAST(parent_id AS text), '') AS "parent_id",
		COALESCE(body, '') AS "body",
		COALESCE(created_at, now()) AS "created_at",
		COALESCE(updated_at, now()) AS "updated_at",
		COALESCE(CAST(user_id AS text), '') AS "user.id"
	FROM
		post_tree
	`

	limit := pagination.Limit + 1 // 次のページがあるかどうかを判定するために1件多く取得する

	queryArgs := []interface{}{parentPostId, limit}

	var statement string

	if pagination.NextCursor == "" {
		statement = baseQuery + `
			ORDER BY reply_created_at ASC -- リプライの作成日時の昇順で取得する
			LIMIT $2;
		`
	} else {
		statement = baseQuery + `
			AND reply_created_at <= (
				SELECT created_at
				FROM reply_relations
				WHERE post_id = $3)
			ORDER BY reply_created_at ASC
			LIMIT $2;
		`

		queryArgs = append(queryArgs, pagination.NextCursor)
	}

	if err := db.SelectContext(ctx, &ps, statement, queryArgs...); err != nil {
		return nil, nil, err
	}

	posts, pg := handlePagination(ps, pagination)

	return posts, pg, nil
}

func (r *Repository) AddPost(ctx context.Context, db Queryer, p *model.Post) (string, error) {
	p.CreatedAt = r.Clocker.Now()
	p.UpdatedAt = r.Clocker.Now()

	statement := `
		INSERT INTO posts (user_id, body, created_at, updated_at)
		VALUES($1, $2, $3, $4)
		RETURNING id;
	`

	queryArgs := []interface{}{p.User.Id, p.Body, p.CreatedAt, p.UpdatedAt}

	var id string
	err := db.QueryRowxContext(ctx, statement, queryArgs...).Scan(&id)

	if err != nil {
		return "", err
	}

	return id, nil
}

func (r *Repository) AddReplyRelation(ctx context.Context, db Execer, postId, parentId string) error {
	statement := `
			INSERT INTO reply_relations (post_id, parent_id, created_at)
			VALUES($1, $2, $3);
		`

	queryArgs := []interface{}{postId, parentId, r.Clocker.Now()}

	_, err := db.ExecContext(ctx, statement, queryArgs...)

	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) DeletePost(ctx context.Context, db Execer, postId string) error {
	statement := `
		DELETE FROM posts WHERE id = $1;
	`

	_, err := db.ExecContext(ctx, statement, postId)

	if err != nil {
		return err
	}

	return nil
}
