package cafeUser

import (
	"bigfood/internal/helpers"
	"database/sql"
)

var ErrorNoResult = sql.ErrNoRows

type Repository interface {
	Add(*CafeUser, Roles) error
	Get(cafeId, userId helpers.Uuid) (*CafeUser, error)
	GetListByCafeId(cafeId helpers.Uuid) ([]*CafeUser, error)
	GetUserRoles(cafeUserId helpers.Uuid) (Roles, error)
	Update(*CafeUser, Roles) error
}
