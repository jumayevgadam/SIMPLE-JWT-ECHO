package connection

import (
	"context"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

// TxOps is
type TxOps interface {
	Begin(ctx context.Context) (pgx.Tx, error)
	Prepare(ctx context.Context, name, query string) (*pgconn.StatementDescription, error)
	Commit(ctx context.Context) error
	RollBack(ctx context.Context) error
	DB
}

// Transaction is
type Transaction struct {
	Tx pgx.Tx
}

// Get is
func (t *Transaction) Get(ctx context.Context, db Querier, dest interface{}, query string, args ...interface{}) error {
	return pgxscan.Get(ctx, t.Tx, dest, query, args...)
}

// Select is
func (t *Transaction) Select(ctx context.Context, db Querier, dest interface{}, query string, args ...interface{}) error {
	return pgxscan.Select(ctx, t.Tx, dest, query, args...)
}

// QueryRow is
func (t *Transaction) QueryRow(ctx context.Context, query string, args ...interface{}) pgx.Row {
	return t.Tx.QueryRow(ctx, query, args...)
}

// Query is
func (t *Transaction) Query(ctx context.Context, query string, args ...interface{}) (pgx.Rows, error) {
	return t.Tx.Query(ctx, query, args...)
}

// Exec is
func (t *Transaction) Exec(ctx context.Context, query string, args ...interface{}) (pgconn.CommandTag, error) {
	return t.Tx.Exec(ctx, query, args...)
}

// Begin is
func (t *Transaction) Begin(ctx context.Context) (pgx.Tx, error) {
	return t.Tx.Begin(ctx)
}

// Prepare is
func (t *Transaction) Prepare(ctx context.Context, name, query string) (*pgconn.StatementDescription, error) {
	return t.Tx.Conn().Prepare(ctx, name, query)
}

// Commit is
func (t *Transaction) Commit(ctx context.Context) error {
	return t.Tx.Commit(ctx)
}

// RollBack is
func (t *Transaction) RollBack(ctx context.Context) error {
	return t.Tx.Rollback(ctx)
}
