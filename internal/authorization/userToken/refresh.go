package userToken

import (
	"github.com/google/uuid"
)

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
		return nil, err
	}
	return &RefreshToken{&token}, nil
}

func (t *RefreshToken) String() string {
	return t.value.String()
}
