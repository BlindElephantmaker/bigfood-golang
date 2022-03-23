package userToken

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

const table = "user_token"

var queryAdd = fmt.Sprintf("INSERT INTO %s (user_id, refresh_token, expires_at) VALUES ($1, $2, $3)", table)

type RepositoryPSQL struct {
	db *sqlx.DB
}

func NewRepositoryPSQL(db *sqlx.DB) *RepositoryPSQL {
	return &RepositoryPSQL{db: db}
}

func (r *RepositoryPSQL) Add(token *UserToken) error {
	_, err := r.db.Exec(queryAdd, token.UserId, token.Refresh, token.ExpiresAt)

	return err
}

func (r *RepositoryPSQL) Get(refreshToken RefreshToken) (*UserToken, error) {
	var userToken UserToken
	query := fmt.Sprintf("SELECT refresh_token, user_id, expires_at FROM %s WHERE refresh_token = $1", table)
	err := r.db.Get(&userToken, query, refreshToken)
	if err != nil {
		return nil, err
	}

	return &userToken, nil
}

func (r *RepositoryPSQL) Refresh(newToken, oldToken *UserToken) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec(queryAdd, newToken.UserId, newToken.Refresh, newToken.ExpiresAt)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	query := fmt.Sprintf("DELETE FROM %s WHERE refresh_token = $1", table)
	_, err = tx.Exec(query, oldToken.Refresh)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	return tx.Commit()
}
