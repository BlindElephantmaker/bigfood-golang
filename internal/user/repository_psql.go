package user

import (
	"bigfood/internal/helpers"
	"fmt"
	"github.com/jmoiron/sqlx"
)

const table = "users"

type userPSQL struct {
	Id    string
	Name  string
	Phone string
}

func (u *userPSQL) castToUser() (*User, error) {
	return Parse(u.Id, u.Name, u.Phone)
}

type RepositoryPSQL struct {
	db *sqlx.DB
}

func NewRepositoryPSQL(db *sqlx.DB) *RepositoryPSQL {
	return &RepositoryPSQL{db: db}
}

func (repo *RepositoryPSQL) Add(u *User) error {
	query := fmt.Sprintf("INSERT INTO %s (id, name, phone) VALUES ($1, $2, $3)", table)
	_, err := repo.db.Exec(query, u.Id, u.Name, u.Phone)

	return err
}

func (repo *RepositoryPSQL) Get(id helpers.Uuid) (*User, error) {
	var userPSQL userPSQL
	query := fmt.Sprintf("SELECT id, name, phone FROM %s WHERE id = $1", table)
	err := repo.db.Get(&userPSQL, query, id)
	if err != nil {
		return nil, err
	}

	return userPSQL.castToUser()
}

func (repo *RepositoryPSQL) Update(u *User) error {
	query := fmt.Sprintf("UPDATE %s SET name = :name WHERE id = :id", table)
	_, err := repo.db.NamedExec(query, map[string]interface{}{
		"id":   u.Id,
		"name": u.Name,
	})

	return err
}

func (repo *RepositoryPSQL) GetByPhone(phone Phone) (*User, error) {
	var userPSQL userPSQL
	query := fmt.Sprintf("SELECT id, name, phone FROM %s WHERE phone = $1", table)
	err := repo.db.Get(&userPSQL, query, phone)
	if err != nil {
		return nil, err
	}

	return userPSQL.castToUser()
}

func (repo *RepositoryPSQL) IsExistByPhone(phone Phone) (bool, error) {
	isExist := false
	query := fmt.Sprintf("SELECT EXISTS(SELECT id, name, phone FROM %s WHERE phone = $1)", table)
	row := repo.db.QueryRow(query, phone)
	err := row.Scan(&isExist)

	return isExist, err
}
