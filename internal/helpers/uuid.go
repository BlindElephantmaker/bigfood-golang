package helpers

import googleUuid "github.com/google/uuid"

type Uuid string

func UuidGenerate() Uuid {
	return Uuid(googleUuid.New().String())
}

func UuidParse(value string) (Uuid, error) {
	uuid, err := googleUuid.Parse(value)
	if err != nil {
		return "", err
	}
	return Uuid(uuid.String()), nil
}
