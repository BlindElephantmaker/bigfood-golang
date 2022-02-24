package user

import (
	"bigfood/internal/helpers"
)

type User struct {
	Id    helpers.Uuid
	Name  Name
	Phone Phone
}

func New(phone Phone) *User {
	return &User{
		Id:    helpers.UuidGenerate(),
		Name:  NewName(),
		Phone: phone,
	}
}

func Parse(idValue, nameValue, phoneValue string) (*User, error) {
	id, err := helpers.UuidParse(idValue)
	if err != nil {
		return nil, err
	}
	name, err := ParseName(nameValue)
	if err != nil {
		return nil, err
	}
	phone, err := NewPhone(phoneValue)
	if err != nil {
		return nil, err
	}

	return &User{
		Id:    id,
		Name:  name,
		Phone: phone,
	}, nil
}

func (u *User) IsNew() bool {
	return u.Name == ""
}
