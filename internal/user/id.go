package user

import (
	"bigfood/internal/helpers"
	"encoding/json"
	"errors"
)

type Id helpers.Uuid

var errorUserIdIsInvalidFormat = errors.New("user id is invalid format")

func (i *Id) UnmarshalJSON(data []byte) error {
	var value string
	err := json.Unmarshal(data, &value)
	if err != nil {
		return err
	}

	uuid, err := helpers.ParseUuid(value)
	if err != nil {
		return errorUserIdIsInvalidFormat
	}

	*i = Id(uuid)
	return nil
}
