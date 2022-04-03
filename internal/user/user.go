package user

import "bigfood/internal/helpers"

type User struct {
	Id    Id            `db:"id"`
	Name  Name          `db:"name"`
	Phone helpers.Phone `db:"phone"`
}

func New(phone helpers.Phone) *User {
	return &User{
		Id:    newId(),
		Name:  NewName(),
		Phone: phone,
	}
}

func (u *User) IsNew() bool {
	return u.Name == ""
}
