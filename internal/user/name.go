package user

import "errors"

type Name struct {
	value string
}

func NewName(username string) (*Name, error) {
	if len(username) == 0 {
		return &Name{username}, nil
	}
	if len(username) < 3 {
		return nil, errors.New("username is too short")
	}
	if len(username) > 64 {
		return nil, errors.New("username is too long")
	}

	return &Name{username}, nil
}

func (n *Name) String() string {
	return n.value
}
