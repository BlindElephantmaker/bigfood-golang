package cafeUser

import (
	"bigfood/internal/cafe"
	"bigfood/internal/user"
	"database/sql"
	"time"
)

var ErrorNoResult = sql.ErrNoRows

type Repository interface {
	Add(*CafeUser, Roles) error
	AddTx(tx *sql.Tx, cafeUser *CafeUser, createAt time.Time, roles Roles) error
	Get(Id) (*CafeUser, error)
	GetByCafe(cafe.Id) ([]*CafeUser, error)
	GetByCafeAndUser(cafe.Id, user.Id) (*CafeUser, error)
	GetUserRoles(Id) (*Roles, error)
	Update(*CafeUser, Roles) error
	Delete(Id) error
}
