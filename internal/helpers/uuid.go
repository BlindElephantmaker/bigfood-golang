package helpers

import (
	"encoding/json"
	"github.com/google/uuid"
)

type Uuid string

func NewUuid() Uuid {
	return Uuid(uuid.New().String())
}

func ParseUuid(data string) (Uuid, error) {
	value, err := uuid.Parse(data)
	if err != nil {
		return "", err
	}
	return Uuid(value.String()), nil
}

func UnmarshalUuid(data []byte) (*Uuid, error) {
	if string(data) == "null" {
		return nil, nil
	}

	var value string
	err := json.Unmarshal(data, &value)
	if err != nil {
		return nil, err
	}

	parsedUuid, err := ParseUuid(value)
	return &parsedUuid, err
}
