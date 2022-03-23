package user

type Repository interface {
	Add(*User) error
	Get(Id) (*User, error)
	Update(*User) error
	GetByPhone(Phone) (*User, error)
	IsExistByPhone(Phone) (bool, error)
}
