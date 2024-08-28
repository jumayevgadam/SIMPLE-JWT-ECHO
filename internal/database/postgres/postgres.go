package postgres

import (
	"SIMPLE-JWT-ECHO/internal/connection"
	"SIMPLE-JWT-ECHO/internal/database"
	"SIMPLE-JWT-ECHO/internal/users"
	userRepository "SIMPLE-JWT-ECHO/internal/users/repository"
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/jackc/pgx/v5"
)

var _ database.DataStore = (*DataStore)(nil)

// DataStore is
type DataStore struct {
	db       connection.DB
	user     users.Repository
	userInit sync.Once
}

// NewDataStore is
func NewDataStore(db connection.DB) database.DataStore {
	return &DataStore{
		db: db,
	}
}

// UsersRepo() is
func (d *DataStore) UsersRepo() users.Repository {
	d.userInit.Do(func() {
		d.user = userRepository.NewUserRepository(d.db)
	})

	return d.user
}

// WithTransaction is
func (d *DataStore) WithTransaction(ctx context.Context, transactionFn database.Transaction) error {
	db, ok := d.db.(connection.DBops)
	if !ok {
		return fmt.Errorf("got error to start transaction")
	}

	tx, err := db.Begin(ctx, pgx.TxOptions{})
	if err != nil {
		return fmt.Errorf("tx: db.Begin: %v", err.Error())
	}

	defer func() {
		if err != nil {
			// RollBack the transaction if an error occured
			if rollBackErr := tx.RollBack(ctx); rollBackErr != nil {
				log.Printf("[postgres][WithTransaction]: failed to rollback transaction: %v", rollBackErr.Error())
			}
		}
	}()

	// transactionalDB is
	transactionalDB := &DataStore{db: tx}
	if err := transactionFn(transactionalDB); err != nil {
		return fmt.Errorf("[postgres][WithTransaction]: transactionFn: %v", err.Error())
	}

	// Commit the transaction if no error occurred during the transactionFn execution
	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("[postgres][WithTransaction]: tx.Commit: %w", err)
	}

	return nil
}
