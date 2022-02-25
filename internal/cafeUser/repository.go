package cafeUser

import (
	"bigfood/internal/cafeUser/role"
	"bigfood/internal/helpers"
	"database/sql"
)

var ErrorNoResult = sql.ErrNoRows

type Repository interface {
	Add(*User) error
	Get(cafeId, userId helpers.Uuid) (*User, error)
	Update(*User) error
	GetUserPermissions(userId helpers.Uuid) (*role.Permissions, error)
}
