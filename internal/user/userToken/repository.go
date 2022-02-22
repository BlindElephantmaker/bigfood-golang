package userToken

import "github.com/google/uuid"

type Repository interface {
	Add(*UserToken) error
	Get(*RefreshToken) (*UserToken, error)
	Delete(token *RefreshToken, userId *uuid.UUID) error
	Refresh(newToken, oldToken *UserToken) error
}
