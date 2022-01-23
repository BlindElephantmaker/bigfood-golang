package user

type Repository interface {
	Add(*User) error
	GetByPhone(*Phone) (*User, error)
	IsExistByPhone(*Phone) (bool, error)
}
