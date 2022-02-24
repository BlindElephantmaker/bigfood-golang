package cafeUser

import (
	"bigfood/internal/cafe/cafeUser/role"
	"github.com/google/uuid"
)

type User struct {
	Id      *uuid.UUID
	CafeId  *uuid.UUID
	Comment Comment
	UserId  *uuid.UUID
	Roles   role.Roles
}

func NewCafeUser(cafeId, userId *uuid.UUID) *User {
	id := uuid.New()
	return &User{
		Id:      &id,
		CafeId:  cafeId,
		UserId:  userId,
		Comment: newComment(),
		Roles: role.Roles{
			role.Owner,
			role.Admin,
			role.Hostess,
		},
	}
}
