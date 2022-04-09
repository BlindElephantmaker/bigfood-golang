package user

import (
	"bigfood/internal/helpers"
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type RepositoryPsql struct {
	db *sqlx.DB
}

func NewRepositoryPsql(db *sqlx.DB) *RepositoryPsql {
	return &RepositoryPsql{db: db}
}

func (r *RepositoryPsql) Add(user *User) error {
	query := fmt.Sprintf("INSERT INTO %s (id, name, phone) VALUES ($1, $2, $3)", table)
	_, err := r.db.Exec(query, user.Id, user.Name, user.Phone)

	return err
}

func (r *RepositoryPsql) Get(userId Id) (*User, error) {
	var user User
	query := fmt.Sprintf("SELECT id, name, phone FROM %s WHERE id = $1", table)
	err := r.db.Get(&user, query, userId)
	if err == sql.ErrNoRows {
		return nil, NotExist
	}
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *RepositoryPsql) GetByPhone(phone helpers.Phone) (*User, error) {
	var user User
	query := fmt.Sprintf("SELECT id, name, phone FROM %s WHERE phone = $1", table)
	err := r.db.Get(&user, query, phone)
	if err == sql.ErrNoRows {
		return nil, NotExist
	}
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *RepositoryPsql) Update(user *User) error {
	query := fmt.Sprintf("UPDATE %s SET name = :name WHERE id = :id", table)
	_, err := r.db.NamedExec(query, map[string]interface{}{
		"id":   user.Id,
		"name": user.Name,
	})

	return err
}
