package cafeUser

import (
	"bigfood/internal/helpers"
	"encoding/json"
	"errors"
)

type Id helpers.Uuid

var errorCafeUserIdIsInvalidFormat = errors.New("cafe user id is invalid format")

func newId() Id {
	return Id(helpers.NewUuid())
}

func (i *Id) UnmarshalJSON(data []byte) error {
	var value string
	err := json.Unmarshal(data, &value)
	if err != nil {
		return err
	}

	uuid, err := helpers.ParseUuid(value)
	if err != nil {
		return errorCafeUserIdIsInvalidFormat
	}

	*i = Id(uuid)
	return nil
}
