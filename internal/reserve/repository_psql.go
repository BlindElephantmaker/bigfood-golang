package reserve

import (
	"bigfood/internal/helpers"
	"bigfood/internal/table"
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"time"
)

const PsqlReserve = "reserve"

type RepositoryPsql struct {
	db *sqlx.DB
}

func NewRepositoryPsql(db *sqlx.DB) *RepositoryPsql {
	return &RepositoryPsql{db: db}
}

func (r *RepositoryPsql) Get(reserveId Id) (*Reserve, error) {
	var reserve Reserve
	query := fmt.Sprintf(`
SELECT id
     , table_id
     , contact_id
     , comment
     , guest_count
     , from_date
     , until_date
     , deleted_at
FROM %s
WHERE id = $1
`, PsqlReserve)
	err := r.db.Get(&reserve, query, reserveId)
	if err == sql.ErrNoRows {
		return nil, notExist
	}
	return &reserve, err
}

func (r *RepositoryPsql) GetActualByTableId(tableId table.Id) ([]*Reserve, error) {
	conditions := "table_id = $1 AND until_date > $2 AND deleted_at IS NULL ORDER BY from_date"
	return r.getList(conditions, tableId, helpers.NowTime())
}

func (r *RepositoryPsql) GetDeletedByTableId(tableId table.Id) ([]*Reserve, error) {
	conditions := "table_id = $1 AND until_date > $2 AND deleted_at IS NOT NULL ORDER BY from_date"
	return r.getList(conditions, tableId, helpers.NowTime())
}

func (r *RepositoryPsql) GetHistoryByTableId(tableId table.Id, limit int, offset int) ([]*Reserve, error) {
	conditions := "table_id = $1 AND until_date < $2 ORDER BY until_date DESC LIMIT $3 OFFSET $4"
	return r.getList(conditions, tableId, helpers.NowTime(), limit, limit*offset)
}

func (r *RepositoryPsql) getList(conditions string, args ...interface{}) ([]*Reserve, error) {
	var reserves []*Reserve
	query := fmt.Sprintf(`
SELECT id
     , table_id
     , contact_id
     , comment
     , guest_count
     , from_date
     , until_date
FROM %s
WHERE %s
`, PsqlReserve, conditions)

	if err := r.db.Select(&reserves, query, args...); err != nil {
		return nil, err
	}

	if len(reserves) == 0 {
		return []*Reserve{}, nil
	}

	return reserves, nil
}

func (r *RepositoryPsql) Add(reserve *Reserve, createdAt time.Time) error {
	if err := r.checkAvailability(reserve); err != nil {
		return err
	}

	query := fmt.Sprintf(`
INSERT INTO %s (id, table_id, contact_id, comment, guest_count, from_date, until_date, created_at)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
`, PsqlReserve)

	_, err := r.db.Exec(
		query,
		reserve.Id,
		reserve.TableId,
		reserve.ContactId,
		reserve.Comment,
		reserve.GuestCount,
		reserve.FromDate,
		reserve.UntilDate,
		createdAt,
	)

	return err
}

func (r *RepositoryPsql) Update(reserve *Reserve) error {
	if err := r.checkAvailability(reserve); err != nil {
		return err
	}

	query := fmt.Sprintf(`
UPDATE %s
SET table_id    = :table_id
  , contact_id  = :contact_id
  , comment     = :comment
  , guest_count = :guest_count
  , from_date   = :from_date
  , until_date  = :until_date
WHERE id = :id
`, PsqlReserve)

	result, err := r.db.NamedExec(query, map[string]interface{}{
		"id":          reserve.Id,
		"table_id":    reserve.TableId,
		"contact_id":  reserve.ContactId,
		"comment":     reserve.Comment,
		"guest_count": reserve.GuestCount,
		"from_date":   reserve.FromDate,
		"until_date":  reserve.UntilDate,
	})
	if err != nil {
		return err
	}
	if count, _ := result.RowsAffected(); count == 0 {
		return notExist
	}

	return nil
}

func (r *RepositoryPsql) Delete(reserveId Id) error {
	now := helpers.NowTime()
	query := fmt.Sprintf("UPDATE %s SET deleted_at = :deleted_at WHERE deleted_at IS NULL AND id = :id", PsqlReserve)
	_, err := r.db.NamedExec(query, map[string]interface{}{
		"id":         reserveId,
		"deleted_at": now,
	})

	return err
}

func (r *RepositoryPsql) Undelete(reserveId Id) error {
	query := fmt.Sprintf("UPDATE %s SET deleted_at = NULL WHERE id = :id", PsqlReserve)
	_, err := r.db.NamedExec(query, map[string]interface{}{
		"id": reserveId,
	})

	return err
}

func (r *RepositoryPsql) checkAvailability(reserve *Reserve) error {
	query := `
SELECT id
FROM reserve
WHERE table_id = :table_id
  AND id != :reserve_id
  AND :from_date < until_date
  AND :until_date > from_date
  AND deleted_at IS NULL
`
	rows, err := r.db.NamedQuery(query, map[string]interface{}{
		"table_id":   reserve.TableId,
		"reserve_id": reserve.Id,
		"from_date":  reserve.FromDate,
		"until_date": reserve.UntilDate,
	})
	if err != nil {
		return err
	}
	if rows.Next() {
		return errorReservedTimeIsBusy
	}

	return nil
}
