package helpers

import (
	"encoding/json"
	googleUuid "github.com/google/uuid"
)

type Uuid string

var errorInvalidUuid = ErrorBadRequest("UUID is invalid")

func NewUuid() Uuid {
	return Uuid(googleUuid.New().String())
}

func (u *Uuid) UnmarshalJSON(data []byte) error {
	var value string
	err := json.Unmarshal(data, &value)
	if err != nil {
		return err
	}

	uuid, err := ParseUuid(value)
	if err != nil {
		return err
	}

	*u = uuid
	return nil
}

func ParseUuid(value string) (Uuid, error) {
	uuid, err := googleUuid.Parse(value)
	if err != nil {
		return "", errorInvalidUuid
	}
	return Uuid(uuid.String()), nil
}
