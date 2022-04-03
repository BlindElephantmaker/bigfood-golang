package cafeUser

import (
	"bigfood/internal/helpers"
	"errors"
)

type Id helpers.Uuid

var errorCafeUserIdIsInvalidFormat = errors.New("cafe user id is invalid format")

func newId() Id {
	return Id(helpers.NewUuid())
}

func (i *Id) UnmarshalJSON(data []byte) error {
	uuid, err := helpers.UnmarshalUuid(data)
	if err != nil {
		return errorCafeUserIdIsInvalidFormat
	}
	*i = Id(*uuid)
	return nil
}
