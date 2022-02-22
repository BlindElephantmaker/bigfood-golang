package cafeUser

import (
	"bigfood/internal/cafe/cafeUser/userRole"
	"github.com/google/uuid"
)

type Repository interface {
	GetUserPermissions(userId *uuid.UUID) (*userRole.Permissions, error)
}
