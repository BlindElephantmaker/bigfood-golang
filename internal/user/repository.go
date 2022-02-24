package user

import (
	"bigfood/internal/helpers"
)

type Repository interface {
	Add(*User) error
	Get(helpers.Uuid) (*User, error)
	Update(*User) error
	GetByPhone(Phone) (*User, error)
	IsExistByPhone(Phone) (bool, error)
}
