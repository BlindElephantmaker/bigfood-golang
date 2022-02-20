package user

import "github.com/google/uuid"

type Repository interface {
	Add(*User) error
	Get(*uuid.UUID) (*User, error)
	Update(*User) error
	GetByPhone(*Phone) (*User, error)
	IsExistByPhone(*Phone) (bool, error)
}
