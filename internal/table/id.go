package table

import (
	"bigfood/internal/helpers"
	"encoding/json"
	"errors"
)

type Id helpers.Uuid

var errorTableIdIsInvalidFormat = errors.New("table id is invalid format")

func NewId() Id {
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
		return errorTableIdIsInvalidFormat
	}

	*i = Id(uuid)
	return nil
}
