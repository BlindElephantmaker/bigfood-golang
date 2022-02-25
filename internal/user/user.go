package user

import (
	"bigfood/internal/helpers"
)

type User struct {
	Id    helpers.Uuid `db:"id"`
	Name  Name         `db:"name"`
	Phone Phone        `db:"phone"`
}

func New(phone Phone) *User {
	return &User{
		Id:    helpers.UuidGenerate(),
		Name:  NewName(),
		Phone: phone,
	}
}

func (u *User) IsNew() bool {
	return u.Name == ""
}
