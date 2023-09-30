package store

import (
	"context"

	"github.com/kngnkg/tunetrail/restapi/model"
)

// ユーザーのフォローを追加する
func (r *Repository) AddFollow(ctx context.Context, db Execer, userId, followeeUserId model.UserID) error {
	createdAt := r.Clocker.Now()
	updatedAt := r.Clocker.Now()

	query := `
		INSERT INTO follows (user_id, followee_id, created_at, updated_at)
		SELECT u1.id, u2.id, $3, $4
		FROM users u1
		INNER JOIN users u2 ON u1.id = $1 AND u2.id = $2
		WHERE u1.is_deleted = false
		AND u2.is_deleted = false;
	`

	_, err := db.ExecContext(ctx, query, userId, followeeUserId, createdAt, updatedAt)
	if err != nil {
		return err
	}
	return nil
}

// ユーザーのフォローを削除する
func (r *Repository) DeleteFollow(ctx context.Context, db Execer, userId, followeeUserId model.UserID) error {
	// 論理削除されていないユーザーかどうかをusersテーブルから確認してからフォローを削除する
	query := `
		DELETE FROM follows
		WHERE user_id = (SELECT id FROM users WHERE id = $1 AND is_deleted = false)
		AND followee_id = (SELECT id FROM users WHERE id = $2 AND is_deleted = false);
	`

	_, err := db.ExecContext(ctx, query, userId, followeeUserId)
	if err != nil {
		return err
	}
	return nil
}
