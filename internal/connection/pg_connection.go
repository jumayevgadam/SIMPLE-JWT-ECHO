package connection

import (
	"context"
	"log"
	"os"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Now we'll write decorators for performing DB operations

// DBops is
var _ DB = (*Database)(nil)

// Querier is
type Querier interface {
	Query(ctx context.Context, query string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, query string, args ...interface{}) pgx.Row
}

// Querier
var (
	_ Querier = &pgxpool.Pool{}
	_ Querier = &pgxpool.Conn{}
)

// DB interface for general database operations
type DB interface {
	Get(ctx context.Context, db Querier, dest interface{}, query string, args ...interface{}) error
	Select(ctx context.Context, db Querier, dest interface{}, query string, args ...interface{}) error
	QueryRow(ctx context.Context, query string, args ...interface{}) pgx.Row
	Query(ctx context.Context, query string, args ...interface{}) (pgx.Rows, error)
	Exec(ctx context.Context, query string, args ...interface{}) (pgconn.CommandTag, error)
}

// DBops interface for general database operations with Transaction
type DBops interface {
	DB
	Begin(ctx context.Context, opts pgx.TxOptions) (TxOps, error)
	Close() error
}

// Database struct implementing the DBops interface
type Database struct {
	db *pgxpool.Pool
}

// GetDBClient is
func GetDBClient(ctx context.Context) (*Database, error) {
	db, err := pgxpool.New(ctx, os.Getenv("connString"))
	if err != nil {
		log.Printf("error in dbConnection")
	}

	err = db.Ping(ctx)
	if err != nil {
		log.Printf("error in ping")
	}

	return &Database{db: db}, nil
}

// Get is
func (d *Database) Get(ctx context.Context, db Querier, dest interface{}, query string, args ...interface{}) error {
	return pgxscan.Get(ctx, db, dest, query, args...)
}

// Select is
func (d *Database) Select(ctx context.Context, db Querier, dest interface{}, query string, args ...interface{}) error {
	return pgxscan.Select(ctx, db, dest, query, args...)
}

// QueryRow is
func (d *Database) QueryRow(ctx context.Context, query string, args ...interface{}) pgx.Row {
	return d.db.QueryRow(ctx, query, args...)
}

// Query is
func (d *Database) Query(ctx context.Context, query string, args ...interface{}) (pgx.Rows, error) {
	return d.db.Query(ctx, query, args...)
}

// Exec is
func (d *Database) Exec(ctx context.Context, query string, args ...interface{}) (pgconn.CommandTag, error) {
	return d.db.Exec(ctx, query, args...)
}

// Begin is
func (d *Database) Begin(ctx context.Context, txOpts pgx.TxOptions) (TxOps, error) {
	c, err := d.db.Acquire(ctx)
	if err != nil {
		return nil, err
	}

	t, err := c.BeginTx(ctx, txOpts)
	if err != nil {
		c.Release()
		return nil, err
	}

	return &Transaction{Tx: t}, nil
}

// Close is
func (d *Database) Close() error {
	d.db.Close()
	return nil
}
