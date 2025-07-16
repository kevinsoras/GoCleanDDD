package postgres

import (
	"context"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Exec(ctx context.Context, db *pgxpool.Pool, query string, args ...interface{}) (pgconn.CommandTag, error) {
	if tx, ok := GetTx(ctx); ok {
		return tx.Exec(ctx, query, args...)
	}
	return db.Exec(ctx, query, args...)
}
