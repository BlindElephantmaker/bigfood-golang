package cafeUser

import (
	"bigfood/internal/cafe/cafeUser/role"
	"bigfood/internal/helpers"
)

type Repository interface {
	GetUserPermissions(userId helpers.Uuid) (*role.Permissions, error)
}
