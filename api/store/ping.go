package store

import (
	"context"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
)

// PingはRDBとの疎通を確認する
func (r *Repository) Ping(ctx context.Context, db Execer) error {
	xdb, ok := db.(*sqlx.DB)
	if !ok {
		return errors.New("invalid args")
	}
	if err := xdb.PingContext(ctx); err != nil {
		return fmt.Errorf("%w: %w", ErrCannotCommunicateWithDB, err)
	}
	return nil
}
