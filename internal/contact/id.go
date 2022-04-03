package contact

import (
	"bigfood/internal/helpers"
	"errors"
)

type Id helpers.Uuid

var errorContactIdIsInvalidFormat = errors.New("contact id is invalid format")

func newId() Id {
	return Id(helpers.NewUuid())
}

func (i *Id) UnmarshalJSON(data []byte) error {
	uuid, err := helpers.UnmarshalUuid(data)
	if err != nil {
		return errorContactIdIsInvalidFormat
	}
	*i = Id(*uuid)
	return nil
}
