package userToken

import (
	"bigfood/internal/helpers"
)

type Repository interface {
	Add(*UserToken) error
	Get(RefreshToken) (*UserToken, error)
	Delete(token RefreshToken, userId helpers.Uuid) error
	Refresh(newToken, oldToken *UserToken) error
}
