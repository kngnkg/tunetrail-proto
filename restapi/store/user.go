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

// GetUserByUserName はユーザー名からユーザーを取得する
func (r *Repository) GetUserByUserName(ctx context.Context, db Queryer, userName string) (*model.User, error) {
	u := &model.User{}
	query := `SELECT id, user_name, name, icon_url, bio, created_at, updated_at
			FROM users
			WHERE user_name = $1;`

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

// UserExistsByUserName はユーザー名が既に存在するかどうかを返す
func (r *Repository) UserExistsByUserName(ctx context.Context, db Queryer, userName string) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM users WHERE user_name = $1);`

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
			WHERE id = $1;`

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

// DeleteUserByUserName はユーザー名からユーザーを削除する
func (r *Repository) DeleteUserByUserName(ctx context.Context, db Execer, userName string) error {
	query := `DELETE FROM users WHERE user_name = $1;`

	row, err := db.ExecContext(ctx, query, userName)
	if err != nil {
		return err
	}
	// 削除された行数を取得
	affected, err := row.RowsAffected()
	if err != nil {
		return err
	}
	// 削除された行数が0の場合はエラーを返す
	if affected == 0 {
		return ErrUserNotFound
	}
	return nil
}
