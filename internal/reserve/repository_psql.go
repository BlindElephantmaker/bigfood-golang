package reserve

import (
	"bigfood/internal/helpers"
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"time"
)

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
`, tableReserve)
	err := r.db.Get(&reserve, query, reserveId)
	if err == sql.ErrNoRows {
		return nil, notExist
	}
	return &reserve, err
}

func (r *RepositoryPsql) Add(reserve *Reserve, createdAt time.Time) error {
	if err := r.checkAvailability(reserve); err != nil {
		return err
	}

	query := fmt.Sprintf(`
INSERT INTO %s (id, table_id, contact_id, comment, guest_count, from_date, until_date, created_at)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
`, tableReserve)

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

func (r *RepositoryPsql) Delete(reserveId Id) error {
	now := helpers.NowTime()
	query := fmt.Sprintf("UPDATE %s SET deleted_at = :deleted_at WHERE deleted_at IS NULL AND id = :id", tableReserve)
	_, err := r.db.NamedExec(query, map[string]interface{}{
		"id":         reserveId,
		"deleted_at": now,
	})

	return err
}

func (r *RepositoryPsql) Undelete(reserveId Id) error {
	query := fmt.Sprintf("UPDATE %s SET deleted_at = NULL WHERE id = :id", tableReserve)
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
  AND :from_date < until_date
  AND :until_date > from_date
`
	rows, err := r.db.NamedQuery(query, map[string]interface{}{
		"table_id":   reserve.TableId,
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
