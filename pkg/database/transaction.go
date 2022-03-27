package database

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
)

type TransactionFactory struct {
	db *sqlx.DB
}

func (s *TransactionFactory) Begin() (*sql.Tx, error) {
	return s.db.Begin()
}

func NewTransactionFactory(db *sqlx.DB) *TransactionFactory {
	return &TransactionFactory{db: db}
}
