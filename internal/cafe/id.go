package cafe

import (
	"bigfood/internal/helpers"
	"errors"
)

type Id helpers.Uuid

var errorCafeIdIsInvalidFormat = errors.New("cafe id is invalid format")

func newId() Id {
	return Id(helpers.NewUuid())
}

func (i *Id) UnmarshalJSON(data []byte) error {
	uuid, err := helpers.UnmarshalUuid(data)
	if err != nil {
		return errorCafeIdIsInvalidFormat
	}
	*i = Id(*uuid)
	return nil
}
