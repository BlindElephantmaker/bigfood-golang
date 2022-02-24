package userToken

import (
	"bigfood/internal/helpers"
	"errors"
)

var ErrorInvalidRefreshTokenFormat = errors.New("invalid refresh token format")

type RefreshToken helpers.Uuid

func NewRefresh() RefreshToken {
	return RefreshToken(helpers.UuidGenerate())
}

func ParseRefresh(tokenValue string) (RefreshToken, error) {
	token, err := helpers.UuidParse(tokenValue)
	if err != nil {
		return "", ErrorInvalidRefreshTokenFormat
	}
	return RefreshToken(token), nil
}
