package reserve

import (
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
