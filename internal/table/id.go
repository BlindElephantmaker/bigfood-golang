package table

import (
	"bigfood/internal/helpers"
	"errors"
)

type Id helpers.Uuid

var errorTableIdIsInvalidFormat = errors.New("table id is invalid format")

func NewId() Id {
	return Id(helpers.NewUuid())
}

func (i *Id) UnmarshalJSON(data []byte) error {
	uuid, err := helpers.UnmarshalUuid(data)
	if err != nil {
		return errorTableIdIsInvalidFormat
	}
	*i = Id(*uuid)
	return nil
}

func ParseId(value string) (Id, error) {
	uuid, err := helpers.ParseUuid(value)
	if err != nil {
		return "", errorTableIdIsInvalidFormat
	}

	return Id(uuid), nil
}
