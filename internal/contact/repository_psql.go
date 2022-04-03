package contact

import (
	"bigfood/internal/helpers"
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"time"
)

type RepositoryPsql struct {
	db *sqlx.DB
}

func NewRepositoryPsql(db *sqlx.DB) *RepositoryPsql {
	return &RepositoryPsql{db: db}
}

func (r *RepositoryPsql) Add(contact Contact, createdAt time.Time) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	query := fmt.Sprintf(`INSERT INTO %s (id, type, created_at) VALUES ($1, $2, $3)`, tableContact)
	if _, err := tx.Exec(query, contact.GetId(), contact.GetType(), createdAt); err != nil {
		_ = tx.Rollback()
		return err
	}

	// todo: add other types
	if err := addPhoneTx(tx, contact.(*Phone)); err != nil {
		_ = tx.Rollback()
		return err
	}

	return tx.Commit()
}

func addPhoneTx(tx *sql.Tx, phone *Phone) error {
	query := fmt.Sprintf(`INSERT INTO %s (contact_id, phone) VALUES ($1, $2)`, tableContactPhone)
	if _, err := tx.Exec(query, phone.GetId(), phone.Phone); err != nil {
		return err
	}
	return nil
}

func (r *RepositoryPsql) GetByPhone(phone *helpers.Phone) (*Phone, error) {
	query := fmt.Sprintf("SELECT contact_id, phone FROM %s WHERE phone = $1", tableContactPhone)

	var contactPhone Phone
	err := r.db.Get(&contactPhone, query, phone)
	if err == sql.ErrNoRows {
		return nil, NotExist
	}
	return &contactPhone, nil
}
