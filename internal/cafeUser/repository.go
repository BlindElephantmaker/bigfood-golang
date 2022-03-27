package cafeUser

import (
	"bigfood/internal/cafe"
	"bigfood/internal/helpers"
	"bigfood/internal/user"
	"database/sql"
	"time"
)

var ErrorNoResult = sql.ErrNoRows

type Repository interface {
	Add(*CafeUser, Roles) error
	AddTx(tx *sql.Tx, cafeUser *CafeUser, createAt time.Time, roles Roles) error
	Get(cafeUserId helpers.Uuid) (*CafeUser, error)
	GetByCafe(cafe.Id) ([]*CafeUser, error)
	GetByCafeAndUser(cafe.Id, user.Id) (*CafeUser, error)
	GetUserRoles(cafeUserId helpers.Uuid) (*Roles, error)
	Update(*CafeUser, Roles) error
	Delete(cafeUserId helpers.Uuid) error
}
