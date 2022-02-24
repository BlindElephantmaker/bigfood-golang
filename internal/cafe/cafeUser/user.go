package cafeUser

import (
	"bigfood/internal/cafe/cafeUser/role"
	"bigfood/internal/helpers"
)

type User struct {
	Id      helpers.Uuid
	CafeId  helpers.Uuid
	Comment Comment
	UserId  helpers.Uuid
	Roles   role.Roles
}

func NewCafeUser(cafeId, userId helpers.Uuid) *User {
	return &User{
		Id:      helpers.UuidGenerate(),
		CafeId:  cafeId,
		UserId:  userId,
		Comment: NewComment(),
		Roles: role.Roles{
			role.Owner,
			role.Admin,
			role.Hostess,
		},
	}
}
