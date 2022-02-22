package cafeUser

import (
	"bigfood/internal/cafe/cafeUser/userRole"
	"github.com/google/uuid"
)

type User struct {
	Id     *uuid.UUID
	CafeId *uuid.UUID
	UserId *uuid.UUID
	Roles  userRole.Roles
}

func NewCafeUser(cafeId, userId *uuid.UUID) *User {
	id := uuid.New()
	return &User{
		Id:     &id,
		CafeId: cafeId,
		UserId: userId,
		Roles: userRole.Roles{
			userRole.Owner,
			userRole.Admin,
			userRole.Hostess,
		},
	}
}
