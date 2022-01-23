package user

import "errors"

type Name struct {
	value string
}

var (
	ErrorUsernameIsTooShort = errors.New("username is too short")
	ErrorUsernameIsTooLong  = errors.New("username is too long")
)

func NewName(username string) (*Name, error) {
	if len(username) == 0 {
		return &Name{username}, nil
	}
	if len(username) < 3 {
		return nil, ErrorUsernameIsTooShort
	}
	if len(username) > 64 {
		return nil, ErrorUsernameIsTooLong
	}

	return &Name{username}, nil
}

func (n *Name) String() string {
	return n.value
}
