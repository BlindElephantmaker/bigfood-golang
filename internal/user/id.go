package user

import (
	"bigfood/internal/helpers"
	"errors"
)

type Id helpers.Uuid

var errorUserIdIsInvalidFormat = errors.New("user id is invalid format")

func newId() Id {
	return Id(helpers.NewUuid())
}

func ParseId(value string) (Id, error) {
	id, err := helpers.ParseUuid(value)
	if err != nil {
		return "", errorUserIdIsInvalidFormat
	}
	return Id(id), nil
}

func (i *Id) UnmarshalJSON(data []byte) error {
	uuid, err := helpers.UnmarshalUuid(data)
	if err != nil {
		return errorUserIdIsInvalidFormat
	}
	*i = Id(*uuid)
	return nil
}
