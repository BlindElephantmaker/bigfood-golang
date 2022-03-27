package cafeUser

import (
	"bigfood/internal/cafe"
	"bigfood/internal/user"
	"database/sql"
	"time"
)

const (
	table     = "cafe_user"
	roleTable = "cafe_user_role"
)

var ErrorNoResult = sql.ErrNoRows

type Repository interface {
	Add(cafeUser *CafeUser, createAt time.Time, roles Roles) error
	AddTx(tx *sql.Tx, cafeUser *CafeUser, createAt time.Time, roles Roles) error
	Get(Id) (*CafeUser, error)
	GetUserRoles(Id) (*Roles, error)
	GetByCafe(cafe.Id) ([]*CafeUser, error)
	GetByCafeAndUser(cafe.Id, user.Id) (*CafeUser, error)
	Update(*CafeUser, Roles) error
	Delete(Id) error
}
