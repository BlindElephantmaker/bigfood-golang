package user

import "bigfood/internal/helpers"

const table = "users"

var NotExist = helpers.ErrorUnprocessableEntity("user not exist")

type Repository interface {
	Add(*User) error
	Get(Id) (*User, error)
	Update(*User) error
	GetByPhone(helpers.Phone) (*User, error)
}
