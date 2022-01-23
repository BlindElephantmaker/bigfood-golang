package userToken

import (
	"errors"
	"github.com/google/uuid"
)

var ErrorInvalidRefreshTokenFormat = errors.New("invalid refresh token format")

type RefreshToken struct {
	value *uuid.UUID
}

func NewRefresh() *RefreshToken {
	value := uuid.New()
	return &RefreshToken{&value}
}

func ParseRefresh(tokenValue string) (*RefreshToken, error) {
	token, err := uuid.Parse(tokenValue)
	if err != nil {
		return nil, ErrorInvalidRefreshTokenFormat
	}
	return &RefreshToken{&token}, nil
}

func (t *RefreshToken) String() string {
	return t.value.String()
}
