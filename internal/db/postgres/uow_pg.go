package postgres

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PgUnitOfWork struct {
	db *pgxpool.Pool
	tx pgx.Tx
}

func NewPgUnitOfWork(db *pgxpool.Pool) *PgUnitOfWork {
	return &PgUnitOfWork{db: db}
}

func (u *PgUnitOfWork) Begin(ctx context.Context) (context.Context, error) {
	tx, err := u.db.Begin(ctx)
	if err != nil {
		return ctx, err
	}
	u.tx = tx
	return WithTx(ctx, tx), nil
}

func (u *PgUnitOfWork) Commit(ctx context.Context) error {
	return u.tx.Commit(ctx)
}

func (u *PgUnitOfWork) Rollback(ctx context.Context) error {
	return u.tx.Rollback(ctx)
}
