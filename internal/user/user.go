package user

import (
	"github.com/google/uuid"
)

type User struct {
	Id    *uuid.UUID
	Name  *Name
	Phone *Phone
}

func New(id *uuid.UUID, name *Name, phone *Phone) *User {
	return &User{
		Id:    id,
		Name:  name,
		Phone: phone,
	}
}

func Parse(idValue, nameValue, phoneValue string) (*User, error) {
	id, err := uuid.Parse(idValue)
	if err != nil {
		return nil, err
	}
	name, err := NewName(nameValue)
	if err != nil {
		return nil, err
	}
	phone, err := NewPhone(phoneValue)
	if err != nil {
		return nil, err
	}

	return New(&id, name, phone), nil
}

func (u *User) IsNew() bool {
	return u.Name.String() == ""
}
