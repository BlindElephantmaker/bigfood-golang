package userToken

import (
	"bigfood/internal/helpers"
)

type RefreshToken helpers.Uuid

func NewRefresh() RefreshToken {
	return RefreshToken(helpers.NewUuid())
}
