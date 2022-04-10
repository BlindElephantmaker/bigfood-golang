package tableListAvailable

import (
	"bigfood/internal/cafe"
	"bigfood/internal/reserve"
	"bigfood/internal/table"
	"fmt"
	"github.com/jmoiron/sqlx"
	"time"
)

type Repository interface {
	GetListAvailable(cafeId cafe.Id, from, until time.Time) ([]*table.Table, error)
}

type RepositoryPsql struct {
	db *sqlx.DB
}

func NewRepositoryPsql(db *sqlx.DB) *RepositoryPsql {
	return &RepositoryPsql{db: db}
}

func (r *RepositoryPsql) GetListAvailable(cafeId cafe.Id, from, until time.Time) ([]*table.Table, error) {
	// todo: condition copy-pasted from reserve repository
	query := fmt.Sprintf(`
SELECT t.id
     , t.cafe_id
     , t.title
     , t.comment
     , t.seats
FROM %s t
         LEFT JOIN %s r ON t.id = r.table_id
    AND $2 < r.until_date
    AND $3 > r.from_date
    AND r.deleted_at IS NULL
WHERE TRUE
  AND cafe_id = $1
  AND r.id IS NULL
`, table.PsqlTables, reserve.PsqlReserve)

	var tables []*table.Table
	if err := r.db.Select(&tables, query, cafeId, from, until); err != nil {
		return nil, err
	}

	if len(tables) == 0 {
		return []*table.Table{}, nil
	}

	return tables, nil
}
