package user

import (
	"encoding/json"
	"errors"
)

type Name string

var (
	errorUsernameIsTooShort = errors.New("username is too short")
	errorUsernameIsTooLong  = errors.New("username is too long")
)

func NewName() Name {
	return ""
}

func (n *Name) UnmarshalJSON(data []byte) error {
	var value string
	err := json.Unmarshal(data, &value)
	if err != nil {
		return err
	}

	name, err := parseName(value)
	if err != nil {
		return err
	}

	*n = name
	return nil
}

func parseName(username string) (Name, error) {
	if len(username) < 3 {
		return "", errorUsernameIsTooShort
	}
	if len(username) > 64 {
		return "", errorUsernameIsTooLong
	}

	return Name(username), nil
}
