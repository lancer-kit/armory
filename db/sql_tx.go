package db

import "github.com/pkg/errors"

// Transactional is the interface for representing
// a db connector/query builder that support database transactions.
type Transactional interface {
	// Begin starts a database transaction.
	Begin() error
	// Commit commits the transaction.
	Commit() error
	// Rollback aborts the transaction.
	Rollback() error
	// InTx checks is transaction started. Return true if it is a transaction, and false if it is not a transaction
	InTx() bool
}

// Transaction is generic helper method for specific Q's to implement Transaction capabilities
func (conn *SQLConn) Transaction(fn func() error) (err error) {
	if err = conn.Begin(); err != nil {
		return errors.Wrap(err, "failed to begin tx")
	}

	if err = fn(); err == nil {
		if err = conn.Commit(); err == nil {
			return nil
		}
		err = errors.Wrap(err, "failed to commit tx")
	} else {
		err = errors.Wrap(err, "failed to execute statements")
	}

	nErr := conn.Rollback()
	if nErr != nil {
		err = errors.Wrapf(err, "rollback failed also(%s)", nErr.Error())
	}
	return
}

// Begin binds this SQLConn to a new transaction.
func (conn *SQLConn) Begin() error {
	if conn.tx != nil {
		return errors.New("already in transaction")
	}

	tx, err := conn.db.Beginx()
	if err != nil {
		return errors.Wrap(err, "failed to begin tx")
	}

	conn.tx = tx
	return nil
}

// Commit commits the current transaction.
func (conn *SQLConn) Commit() error {
	if conn.tx == nil {
		return errors.New("not in transaction")
	}

	err := conn.tx.Commit()
	if err != nil {
		return err
	}

	conn.tx = nil
	return nil
}

// Rollback rolls back the current transaction
func (conn *SQLConn) Rollback() error {
	if conn.tx == nil {
		return nil
	}

	err := conn.tx.Rollback()
	conn.tx = nil
	return err
}

// InTx checks is transaction started. Return true if it is a transaction, and false if it is not a transaction
func (conn *SQLConn) InTx() bool {
	return conn.tx != nil
}
