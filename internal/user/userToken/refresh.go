package userToken

import (
	"bigfood/internal/helpers"
	"errors"
)

type RefreshToken helpers.Uuid

var errorRefreshTokenIsInvalidFormat = errors.New("refresh token is invalid format")

func newRefresh() RefreshToken {
	return RefreshToken(helpers.NewUuid())
}

func (rt *RefreshToken) UnmarshalJSON(data []byte) error {
	uuid, err := helpers.UnmarshalUuid(data)
	if err != nil {
		return errorRefreshTokenIsInvalidFormat
	}
	*rt = RefreshToken(*uuid)
	return nil
}
