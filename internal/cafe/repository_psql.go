package cafe

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"time"
)

type RepositoryPsql struct {
	*sqlx.DB
}

func NewRepositoryPsql(db *sqlx.DB) *RepositoryPsql {
	return &RepositoryPsql{db}
}

func (r *RepositoryPsql) AddTx(tx *sql.Tx, cafe *Cafe, createdAt time.Time) error {
	query := fmt.Sprintf("INSERT INTO %s (id, created_at) VALUES ($1, $2)", table)
	_, err := tx.Exec(query, cafe.Id, createdAt)
	return err
}
