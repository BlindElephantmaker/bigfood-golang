package userToken

import (
	"bigfood/internal/helpers/timeHelper"
	"github.com/google/uuid"
	"time"
)

const (
	day        = 24 * time.Hour
	accessTTL  = 1 * day
	refreshTTL = 360 * day
)

type UserToken struct {
	UserId    *uuid.UUID
	Access    *AccessToken
	Refresh   *RefreshToken
	ExpiresAt *time.Time
}

func New(userId *uuid.UUID) (*UserToken, error) {
	now := timeHelper.Now()
	expiresAt := now.Add(refreshTTL)

	access, err := NewAccess(userId, &now, accessTTL)
	if err != nil {
		return nil, err
	}

	refresh := NewRefresh()
	if err != nil {
		return nil, err
	}

	return &UserToken{
		UserId:    userId,
		Access:    access,
		Refresh:   refresh,
		ExpiresAt: &expiresAt,
	}, nil
}

func Parse(userIdValue, refreshValue, expiresAtValue string) (*UserToken, error) {
	userId, err := uuid.Parse(userIdValue)
	if err != nil {
		return nil, err
	}
	refresh, err := ParseRefresh(refreshValue)
	if err != nil {
		return nil, err
	}
	expiresAt, err := timeHelper.Parse(expiresAtValue)
	if err != nil {
		return nil, err
	}

	return &UserToken{
		UserId:    &userId,
		Access:    nil,
		Refresh:   refresh,
		ExpiresAt: &expiresAt,
	}, nil
}
