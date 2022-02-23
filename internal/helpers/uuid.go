package helpers

import (
	"errors"
	googleUuid "github.com/google/uuid"
)

type Uuid string

var ErrorInvalidUuid = errors.New("UUID is invalid")

func UuidGenerate() Uuid {
	return Uuid(googleUuid.New().String())
}

func UuidParse(value string) (Uuid, error) {
	uuid, err := googleUuid.Parse(value)
	if err != nil {
		return "", ErrorInvalidUuid
	}
	return Uuid(uuid.String()), nil
}
