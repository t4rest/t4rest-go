package mysql

import (
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

// Tx .
type Tx interface {
	GetTx() *sqlx.Tx
	Commit() error
	Rollback() error
}

// myTx .
type myTx struct {
	tx *sqlx.Tx
}

// NewTx .
func NewTx(tx *sqlx.Tx) Tx {
	return &myTx{tx: tx}
}

// GetTx .
func (myTx *myTx) GetTx() *sqlx.Tx {
	return myTx.tx
}

// Commit .
func (myTx *myTx) Commit() error {
	if myTx.tx == nil {
		return errors.New("tx is not initialized")
	}

	return myTx.tx.Commit()
}

// Rollback .
func (myTx *myTx) Rollback() error {
	if myTx.tx == nil {
		return nil
	}

	return myTx.tx.Rollback()
}
