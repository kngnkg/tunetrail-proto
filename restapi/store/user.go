package store

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/kngnkg/tunetrail/restapi/model"
	"github.com/lib/pq"
)

const (
	// usersテーブルのユニーク制約
	ConstraintUserName = "users_user_name_key"
)

// GetUserByUserId はユーザーIDからユーザーを取得する
func (r *Repository) GetUserByUserId(ctx context.Context, db Queryer, id model.UserID) (*model.User, error) {
	u := &model.User{}
	query := `SELECT id, user_name, name, icon_url, bio, created_at, updated_at
			FROM users
			WHERE is_deleted = false
			AND id = $1;`

	if err := db.GetContext(ctx, u, query, string(id)); err != nil {
		// ユーザーが存在しない場合はエラーをラップして返す
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("%w: %w", ErrUserNotFound, err)
		}
		return nil, err
	}
	// UTCに変換
	u.CreatedAt = u.CreatedAt.UTC()
	u.UpdatedAt = u.UpdatedAt.UTC()
	return u, nil
}

// GetUserByUserName はユーザー名からユーザーを取得する
func (r *Repository) GetUserByUserName(ctx context.Context, db Queryer, userName string) (*model.User, error) {
	u := &model.User{}
	query := `SELECT id, user_name, name, icon_url, bio, created_at, updated_at
			FROM users
			WHERE is_deleted = false
			AND user_name = $1;`

	if err := db.GetContext(ctx, u, query, userName); err != nil {
		// ユーザーが存在しない場合はエラーをラップして返す
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("%w: %w", ErrUserNotFound, err)
		}
		return nil, err
	}
	// UTCに変換
	u.CreatedAt = u.CreatedAt.UTC()
	u.UpdatedAt = u.UpdatedAt.UTC()
	return u, nil
}

// 暫定
// TODO: フォロー情報を取得する処理を別のメソッドに分ける
func (r *Repository) GetUserByUserNameWithFollowInfo(ctx context.Context, db Queryer, userName string, signedInUserId model.UserID) (*model.User, error) {
	u := &model.User{}

	// f1: ログインユーザーがフォローしている
	// f2: ログインユーザーがフォローされている
	query := `
	SELECT
		u.id,
		u.user_name,
		u.name,
		u.icon_url,
		u.bio,
		CASE WHEN f1.user_id IS NULL THEN false ELSE true END AS is_following,
		CASE WHEN f2.followee_id IS NULL THEN false ELSE true END AS is_followed,
		u.created_at,
		u.updated_at
	FROM users u
	LEFT OUTER JOIN follows f1 ON u.id = f1.followee_id AND f1.user_id = $2
	LEFT OUTER JOIN follows f2 ON u.id = f2.user_id AND f2.followee_id = $2
	WHERE u.user_name = $1
	AND u.is_deleted = false
	`

	if err := db.GetContext(ctx, u, query, userName, signedInUserId); err != nil {
		// ユーザーが存在しない場合はエラーをラップして返す
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("%w: %w", ErrUserNotFound, err)
		}
		return nil, err
	}
	// UTCに変換
	u.CreatedAt = u.CreatedAt.UTC()
	u.UpdatedAt = u.UpdatedAt.UTC()
	return u, nil
}

// UserExistsByUserName はユーザー名が既に存在するかどうかを返す
func (r *Repository) UserExistsByUserName(ctx context.Context, db Queryer, userName string) (bool, error) {
	query := `SELECT EXISTS(
			SELECT 1 FROM users
			WHERE is_deleted = false
			AND user_name = $1
			);`

	var exists bool
	err := db.QueryRowxContext(ctx, query, userName).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}

// RegisterUser はユーザーを登録する
func (r *Repository) RegisterUser(ctx context.Context, db Execer, u *model.User) error {
	u.CreatedAt = r.Clocker.Now()
	u.UpdatedAt = r.Clocker.Now()
	query := `INSERT INTO users (id, user_name, name, icon_url, bio, created_at, updated_at)
			VALUES ($1, $2, $3, $4, $5, $6, $7);`

	_, err := db.ExecContext(ctx, query, u.Id, u.UserName, u.Name, u.IconUrl, u.Bio, u.CreatedAt, u.UpdatedAt)
	if err != nil {
		var pqError *pq.Error
		// 重複エラーの場合はエラーをラップして返す
		if errors.As(err, &pqError) && pqError.Code == ErrCodePostgresDuplicate {
			if pqError.Constraint == ConstraintUserName {
				return fmt.Errorf("%w: %w", ErrUserNameAlreadyExists, err)
			}
		}
		return err
	}

	return nil
}

// UpdateUser はユーザーを更新する
func (r *Repository) UpdateUser(ctx context.Context, db Execer, u *model.User) error {
	u.UpdatedAt = r.Clocker.Now()
	query := `UPDATE users
			SET user_name = $2, name = $3, icon_url = $4, bio = $5, updated_at = $6
			WHERE is_deleted = false
			AND id = $1;`

	row, err := db.ExecContext(ctx, query, u.Id, u.UserName, u.Name, u.IconUrl, u.Bio, u.UpdatedAt)
	if err != nil {
		var pqError *pq.Error
		// 重複エラーの場合はエラーをラップして返す
		if errors.As(err, &pqError) && pqError.Code == ErrCodePostgresDuplicate {
			// どの制約が違反されたかでエラーを分ける
			if pqError.Constraint == ConstraintUserName {
				return fmt.Errorf("%w: %w", ErrUserNameAlreadyExists, err)
			}
		}
		return err
	}
	// 更新された行数を取得
	affected, err := row.RowsAffected()
	if err != nil {
		return err
	}
	// 更新された行数が0の場合はエラーを返す
	if affected == 0 {
		return ErrUserNotFound
	}

	// UTCに変換
	u.CreatedAt = u.CreatedAt.UTC()
	u.UpdatedAt = u.UpdatedAt.UTC()
	return nil
}

// DeleteUserByUserName はユーザー名からユーザーを論理削除する
func (r *Repository) DeleteUserByUserName(ctx context.Context, db Execer, userName string) error {
	query := `UPDATE users
			SET is_deleted = true
			WHERE is_deleted = false
			AND user_name = $1;`

	row, err := db.ExecContext(ctx, query, userName)
	if err != nil {
		return err
	}
	// 更新された行数を取得
	affected, err := row.RowsAffected()
	if err != nil {
		return err
	}
	// 更新された行数が0の場合はエラーを返す
	if affected == 0 {
		return ErrUserNotFound
	}
	return nil
}

// ユーザーのフォローを追加する
func (r *Repository) AddFollow(ctx context.Context, db Execer, userName, followeeUserName string) error {
	createdAt := r.Clocker.Now()
	updatedAt := r.Clocker.Now()

	// ユーザーテーブルからフォローするユーザーとフォローされるユーザーのIDを取得してからフォローを追加する
	query := `
	INSERT INTO follows (user_id, followee_id, created_at, updated_at)
	SELECT u1.id, u2.id, $3, $4
	FROM users u1
	INNER JOIN users u2 ON u1.user_name = $1 AND u2.user_name = $2
	WHERE u1.is_deleted = false
	AND u2.is_deleted = false;
	`

	_, err := db.ExecContext(ctx, query, userName, followeeUserName, createdAt, updatedAt)
	if err != nil {
		return err
	}
	return nil
}

// ユーザーのフォローを削除する
func (r *Repository) DeleteFollow(ctx context.Context, db Execer, userName, followeeUserName string) error {
	// 論理削除されていないユーザーかどうかをusersテーブルから確認してからフォローを削除する
	query := `
	DELETE FROM follows
	WHERE user_id = (SELECT id FROM users WHERE user_name = $1 AND is_deleted = false)
	AND followee_id = (SELECT id FROM users WHERE user_name = $2 AND is_deleted = false);
	`

	_, err := db.ExecContext(ctx, query, userName, followeeUserName)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) GetFolloweesByUserId(ctx context.Context, db Queryer, userId model.UserID) ([]*model.User, error) {
	var users []*model.User
	query := `
	SELECT
		u.id,
		u.user_name,
		u.name,
		u.icon_url,
		u.bio,
		u.created_at,
		u.updated_at
	FROM users u
	INNER JOIN follows f ON u.id = f.followee_id AND user_id = $1
	WHERE u.is_deleted = false;
	`

	if err := db.SelectContext(ctx, &users, query, userId); err != nil {
		return nil, err
	}

	// UTCに変換
	for _, u := range users {
		u.CreatedAt = u.CreatedAt.UTC()
		u.UpdatedAt = u.UpdatedAt.UTC()
	}

	return users, nil
}
