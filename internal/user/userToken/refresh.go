package userToken

import (
	"bigfood/internal/helpers"
	"encoding/json"
	"errors"
)

type RefreshToken helpers.Uuid

var errorRefreshTokenIsInvalidFormat = errors.New("refresh token is invalid format")

func NewRefresh() RefreshToken {
	return RefreshToken(helpers.NewUuid())
}

func (rt *RefreshToken) UnmarshalJSON(data []byte) error {
	var value string
	err := json.Unmarshal(data, &value)
	if err != nil {
		return err
	}

	uuid, err := helpers.ParseUuid(value)
	if err != nil {
		return errorRefreshTokenIsInvalidFormat
	}

	*rt = RefreshToken(uuid)
	return nil
}
