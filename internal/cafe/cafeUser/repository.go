package cafeUser

import (
	"bigfood/internal/cafe/cafeUser/role"
	"github.com/google/uuid"
)

type Repository interface {
	GetUserPermissions(userId *uuid.UUID) (*role.Permissions, error)
}
