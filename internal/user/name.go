package user

import "errors"

type Name string

var (
	ErrorUsernameIsTooShort = errors.New("username is too short")
	ErrorUsernameIsTooLong  = errors.New("username is too long")
)

func NewName() Name {
	return ""
}

func ParseName(username string) (Name, error) {
	if len(username) < 3 {
		return "", ErrorUsernameIsTooShort
	}
	if len(username) > 64 {
		return "", ErrorUsernameIsTooLong
	}

	return Name(username), nil
}
