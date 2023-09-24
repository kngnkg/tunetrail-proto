package service

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/kngnkg/tunetrail/restapi/store"
)

type Transactioner interface {
	WithTransaction(ctx context.Context, db store.Beginner, f func(tx *sqlx.Tx) error) error
}
