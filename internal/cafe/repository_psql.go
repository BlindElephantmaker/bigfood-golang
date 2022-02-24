package cafe

import (
	"bigfood/internal/cafe/cafeUser"
	"fmt"
	"github.com/jmoiron/sqlx"
	"time"
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

func (r *RepositoryPSQL) Add(cafe *Cafe, cafeUser *cafeUser.User, createAt *time.Time) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	queryCafe := fmt.Sprintf("INSERT INTO %s (id, created_at) VALUES ($1, $2)", cafeTable)
	result, err := tx.Exec(queryCafe, cafe.Id, createAt)
	fmt.Println(result)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	err = r.CafeUserRepository.AddTx(tx, cafeUser, createAt)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	return tx.Commit()
}
