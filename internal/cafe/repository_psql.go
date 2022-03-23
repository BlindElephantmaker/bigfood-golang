package cafe

import (
	"bigfood/internal/cafeUser"
	"bigfood/internal/helpers"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type RepositoryPSQL struct {
	db                 *sqlx.DB
	CafeUserRepository *cafeUser.RepositoryPSQL
}

func NewRepositoryPSQL(db *sqlx.DB, cafeUserRepository *cafeUser.RepositoryPSQL) *RepositoryPSQL {
	return &RepositoryPSQL{
		db:                 db,
		CafeUserRepository: cafeUserRepository,
	}
}

const (
	cafeTable = "cafe"
)

func (r *RepositoryPSQL) Add(cafe *Cafe, cafeUser *cafeUser.CafeUser, userRoles cafeUser.Roles) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	createAt := helpers.NowTime()
	queryCafe := fmt.Sprintf("INSERT INTO %s (id, created_at) VALUES ($1, $2)", cafeTable)
	if _, err := tx.Exec(queryCafe, cafe.Id, createAt); err != nil {
		_ = tx.Rollback()
		return err
	}

	if err := r.CafeUserRepository.AddTx(tx, cafeUser, userRoles, createAt); err != nil {
		_ = tx.Rollback()
		return err
	}

	return tx.Commit()
}
