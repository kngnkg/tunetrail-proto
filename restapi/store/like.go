package store

import (
	"context"

	"github.com/kngnkg/tunetrail/restapi/model"
)

func (r *Repository) AddLike(ctx context.Context, db Execer, userId model.UserID, postId string) error {
	statement := `
		INSERT INTO likes (post_id, user_id, created_at, updated_at)
		SELECT p.id, u.id, $3, $4
		FROM posts p
		INNER JOIN users u
		ON p.id = $1 AND u.id = $2;
	`

	createdAt := r.Clocker.Now()
	updatedAt := r.Clocker.Now()

	queryArgs := []interface{}{postId, userId, createdAt, updatedAt}

	_, err := db.ExecContext(ctx, statement, queryArgs...)

	if err != nil {
		return err
	}

	return nil
}
