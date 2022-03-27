package user

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

const table = "users"

type RepositoryPsql struct {
	db *sqlx.DB
}

func NewRepositoryPsql(db *sqlx.DB) *RepositoryPsql {
	return &RepositoryPsql{db: db}
}

func (r *RepositoryPsql) Add(u *User) error {
	query := fmt.Sprintf("INSERT INTO %s (id, name, phone) VALUES ($1, $2, $3)", table)
	_, err := r.db.Exec(query, u.Id, u.Name, u.Phone)

	return err
}

func (r *RepositoryPsql) Get(id Id) (*User, error) {
	var user User
	query := fmt.Sprintf("SELECT id, name, phone FROM %s WHERE id = $1", table)
	if err := r.db.Get(&user, query, id); err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *RepositoryPsql) Update(u *User) error {
	query := fmt.Sprintf("UPDATE %s SET name = :name WHERE id = :id", table)
	_, err := r.db.NamedExec(query, map[string]interface{}{
		"id":   u.Id,
		"name": u.Name,
	})

	return err
}

func (r *RepositoryPsql) GetByPhone(phone Phone) (*User, error) {
	var user User
	query := fmt.Sprintf("SELECT id, name, phone FROM %s WHERE phone = $1", table)
	if err := r.db.Get(&user, query, phone); err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *RepositoryPsql) IsExistByPhone(phone Phone) (bool, error) {
	isExist := false
	query := fmt.Sprintf("SELECT EXISTS(SELECT id, name, phone FROM %s WHERE phone = $1)", table)
	row := r.db.QueryRow(query, phone)
	err := row.Scan(&isExist)

	return isExist, err
}
