package table

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"strings"
	"time"
)

type RepositoryPSQL struct {
	db *sqlx.DB
}

func NewRepositoryPSQL(db *sqlx.DB) *RepositoryPSQL {
	return &RepositoryPSQL{db: db}
}

const table = "tables"

func (r *RepositoryPSQL) Add(tables []*Table, createdAt time.Time) error {
	params := []interface{}{}

	for _, table := range tables {
		params = append(params,
			table.Id,
			table.CafeId,
			table.Title,
			table.Comment,
			table.Seats,
			createdAt,
		)
	}

	query := createInsertQuery(len(tables))
	_, err := r.db.Exec(query, params...)
	return err
}

func createInsertQuery(batchSize int) string {
	var buffer []string

	// todo: lol, fix it
	for i := 0; i < batchSize; i++ {
		params := []interface{}{
			i*6 + 1,
			i*6 + 2,
			i*6 + 3,
			i*6 + 4,
			i*6 + 5,
			i*6 + 6,
		}
		sql := fmt.Sprintf("($%v, $%v, $%v, $%v, $%v, $%v)", params...)
		buffer = append(buffer, sql)
	}

	query := fmt.Sprintf("INSERT INTO %s (id, cafe_id, title, comment, seats, created_at) VALUES %s",
		table,
		strings.Join(buffer, ","),
	)

	return query
}
