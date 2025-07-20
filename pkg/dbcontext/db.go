package dbcontext

import (
	"gorm.io/gorm"
	"context"
)

// DB represents a DB connection that can be used to run SQL queries.
type DB struct {
	db *gorm.DB
}

// New returns a new DB connection that wraps the given gorm.DB instance.
func New(db *gorm.DB) *DB {
	return &DB{db}
}

// With returns a Builder that can be used to build and execute SQL queries.
// With will return the transaction if it is found in the given context.
// Otherwise it will return a DB connection associated with the context.
func (db *DB) With(ctx context.Context) *gorm.DB {
	// if tx, ok := ctx.Value(txKey).(*dbx.Tx); ok {
	// 	return tx
	// }
	return db.db.WithContext(ctx)
}