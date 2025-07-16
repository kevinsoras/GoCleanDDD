package postgres

import (
	"context"

	"github.com/jackc/pgx/v5"
)

// Se usa un Context-aware para mantener la transacción en el contexto
type txKeyType struct{}

var txKey = txKeyType{}

func WithTx(ctx context.Context, tx pgx.Tx) context.Context {
	return context.WithValue(ctx, txKey, tx)
}

func GetTx(ctx context.Context) (pgx.Tx, bool) {
	tx, ok := ctx.Value(txKey).(pgx.Tx)
	return tx, ok
}
