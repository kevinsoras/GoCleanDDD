package postgres

import (
	"context"
)

// UnitOfWork define la interface de unidad de trabajo
type UnitOfWork interface {
	Begin(ctx context.Context) (context.Context, error)
	Commit(ctx context.Context) error
	Rollback(ctx context.Context) error
}
