package userToken

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

const table = "user_token"

var queryAdd = fmt.Sprintf("INSERT INTO %s (user_id, refresh_token, expires_at) VALUES ($1, $2, $3)", table)

type userTokenPSQL struct {
	RefreshToken string `db:"refresh_token"`
	UserId       string `db:"user_id"`
	ExpiresAt    string `db:"expires_at"`
}

func (ut *userTokenPSQL) castToUserToken() (*UserToken, error) {
	return Parse(ut.UserId, ut.RefreshToken, ut.ExpiresAt)
}

type RepositoryPSQL struct {
	db *sqlx.DB
}

func NewRepositoryPSQL(db *sqlx.DB) *RepositoryPSQL {
	return &RepositoryPSQL{db: db}
}

func (r *RepositoryPSQL) Add(token *UserToken) error {
	_, err := r.db.Exec(queryAdd, token.UserId.String(), token.Refresh.String(), token.ExpiresAt)

	return err
}

func (r *RepositoryPSQL) Get(refreshToken *RefreshToken) (*UserToken, error) {
	var tokenPSQL userTokenPSQL
	query := fmt.Sprintf("SELECT refresh_token, user_id, expires_at FROM %s WHERE refresh_token = $1", table)
	err := r.db.Get(&tokenPSQL, query, refreshToken.String())
	if err != nil {
		return nil, err
	}

	return tokenPSQL.castToUserToken()
}

func (r *RepositoryPSQL) Refresh(newToken, oldToken *UserToken) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec(queryAdd, newToken.UserId.String(), newToken.Refresh.String(), newToken.ExpiresAt)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	query := fmt.Sprintf("DELETE FROM %s WHERE refresh_token = $1", table)
	_, err = tx.Exec(query, oldToken.Refresh.String())
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	return tx.Commit()
}
