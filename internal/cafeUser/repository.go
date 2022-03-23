package cafeUser

import (
	"bigfood/internal/helpers"
	"bigfood/internal/user"
	"database/sql"
)

var ErrorNoResult = sql.ErrNoRows

type Repository interface {
	Add(*CafeUser, Roles) error
	Get(cafeUserId helpers.Uuid) (*CafeUser, error)
	GetByCafeAndUserIds(cafeId helpers.Uuid, userId user.Id) (*CafeUser, error)
	GetListByCafeId(cafeId helpers.Uuid) ([]*CafeUser, error)
	GetUserRoles(cafeUserId helpers.Uuid) (*Roles, error)
	Update(*CafeUser, Roles) error
	Delete(cafeUserId helpers.Uuid) error
}
