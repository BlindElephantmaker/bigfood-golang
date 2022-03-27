package cafe

import (
	"database/sql"
	"time"
)

const table = "cafe"

type Repository interface {
	AddTx(tx *sql.Tx, cafe *Cafe, createAt time.Time) error
}
