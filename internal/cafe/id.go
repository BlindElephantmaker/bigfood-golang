package cafe

import (
	"bigfood/internal/helpers"
	"encoding/json"
	"errors"
)

type Id helpers.Uuid

var errorCafeIdIsInvalidFormat = errors.New("cafe id is invalid format")

func NewCafeId() Id {
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
		return errorCafeIdIsInvalidFormat
	}

	*i = Id(uuid)
	return nil
}
