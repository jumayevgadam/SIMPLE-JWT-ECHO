package database

import (
	"SIMPLE-JWT-ECHO/internal/users"
	"context"
)

// Transaction is
type Transaction func(db DataStore) error

// DataStore is
type DataStore interface {
	WithTransaction(ctx context.Context, tx Transaction) error
	UsersRepo() users.Repository
}
