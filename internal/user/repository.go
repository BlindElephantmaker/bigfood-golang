package user

const table = "users"

type Repository interface {
	Add(*User) error
	Get(Id) (*User, error)
	Update(*User) error
	GetByPhone(Phone) (*User, error)
	IsExistByPhone(Phone) (bool, error)
}
