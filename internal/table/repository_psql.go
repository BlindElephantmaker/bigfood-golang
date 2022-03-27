package table

import (
	"bigfood/internal/cafe"
	"bigfood/internal/helpers"
	"fmt"
	"github.com/jmoiron/sqlx"
	"strings"
	"time"
)

type RepositoryPsql struct {
	db *sqlx.DB
}

func NewRepositoryPsql(db *sqlx.DB) *RepositoryPsql {
	return &RepositoryPsql{db: db}
}

func (r *RepositoryPsql) AddSlice(tables []*Table, createdAt time.Time) error {
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

func (r *RepositoryPsql) Get(tableId Id) (*Table, error) {
	var t Table
	query := fmt.Sprintf("SELECT id, cafe_id, title, comment, seats FROM %s WHERE id = $1", table)
	err := r.db.Get(&t, query, tableId)
	if err != nil {
		return nil, err
	}

	return &t, nil
}

func (r *RepositoryPsql) GetByCafe(cafeId cafe.Id) ([]*Table, error) {
	var tables []*Table
	query := fmt.Sprintf(
		"SELECT id, cafe_id, title, comment, seats FROM %s WHERE cafe_id = $1 AND deleted_at IS NULL ORDER BY title",
		table,
	)
	err := r.db.Select(&tables, query, cafeId)
	if err != nil {
		return nil, err
	}

	return tables, nil
}

func (r *RepositoryPsql) Update(t *Table) error {
	query := fmt.Sprintf("UPDATE %s SET title = :title, comment = :comment, seats = :seats WHERE id = :id", table)
	_, err := r.db.NamedExec(query, map[string]interface{}{
		"id":      t.Id,
		"title":   t.Title,
		"comment": t.Comment,
		"seats":   t.Seats,
	})

	return err
}

func (r *RepositoryPsql) Delete(tableId Id) error {
	now := helpers.NowTime()
	query := fmt.Sprintf("UPDATE %s SET deleted_at = :deleted_at WHERE deleted_at IS NULL AND id = :id", table)
	_, err := r.db.NamedExec(query, map[string]interface{}{
		"id":         tableId,
		"deleted_at": now,
	})

	return err
}

func (r *RepositoryPsql) DeleteAll(cafeId cafe.Id) error {
	now := helpers.NowTime()
	query := fmt.Sprintf("UPDATE %s SET deleted_at = :deleted_at WHERE deleted_at IS NULL AND cafe_id = :cafe_id", table)
	_, err := r.db.NamedExec(query, map[string]interface{}{
		"cafe_id":    cafeId,
		"deleted_at": now,
	})

	return err
}
