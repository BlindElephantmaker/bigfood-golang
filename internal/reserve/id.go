package reserve

import (
	"bigfood/internal/helpers"
	"errors"
)

type Id helpers.Uuid

var errorReserveIdIsInvalidFormat = errors.New("reserve id is invalid format")

func newId() Id {
	return Id(helpers.NewUuid())
}

func (i *Id) UnmarshalJSON(data []byte) error {
	uuid, err := helpers.UnmarshalUuid(data)
	if err != nil {
		return errorReserveIdIsInvalidFormat
	}
	*i = Id(*uuid)
	return nil
}

func ParseId(value string) (Id, error) {
	uuid, err := helpers.ParseUuid(value)
	if err != nil {
		return "", errorReserveIdIsInvalidFormat
	}

	return Id(uuid), nil
}
