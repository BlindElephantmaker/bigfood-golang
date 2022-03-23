package user

import (
	"bigfood/internal/helpers"
)

type User struct {
	Id    Id    `db:"id"`
	Name  Name  `db:"name"`
	Phone Phone `db:"phone"`
}

func New(phone Phone) *User {
	return &User{
		Id:    Id(helpers.NewUuid()),
		Name:  NewName(),
		Phone: phone,
	}
}

func (u *User) IsNew() bool {
	return u.Name == ""
}
