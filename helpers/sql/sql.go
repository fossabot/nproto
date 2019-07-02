package sqlh

import (
	"context"
	"database/sql"
)

// Queryer abstracts sql.DB/sql.Conn/sql.Tx .
type Queryer interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

// TxPlugin is used to add extra functionalities during transaction life time.
type TxPlugin interface {
	// TxInitialized will be called after a transaction started. If an error returned, the transaction
	// will be rollbacked at once.
	TxInitialized(db *sql.DB, tx *sql.Tx) error
	// TxCommitted will be called only after the transaction has been successfully committed.
	TxCommitted()
	// TxFinalised will be called after the transaction committed/rollbacked.
	TxFinalised()
}

var (
	_ Queryer = (*sql.DB)(nil)
	_ Queryer = (*sql.Conn)(nil)
	_ Queryer = (*sql.Tx)(nil)
)

// WithTx starts a transaction and run fn. If no error is returned, the transaction is committed.
// Otherwise it is rollbacked.
func WithTx(db *sql.DB, fn func(*sql.Tx) error, plugins ...TxPlugin) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	var n int
	defer func() {
		tx.Rollback()
		for _, plugin := range plugins[:n] {
			plugin.TxFinalised()
		}
	}()

	for _, plugin := range plugins {
		if err := plugin.TxInitialized(db, tx); err != nil {
			return err
		}
		n += 1
	}

	err = fn(tx)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	for _, plugin := range plugins {
		plugin.TxCommitted()
	}
	return nil
}
