package userToken

import (
	"bigfood/internal/helpers"
	"bigfood/internal/user"
	"time"
)

const (
	day        = 24 * time.Hour
	accessTTL  = 1 * day
	refreshTTL = 360 * day
)

type UserToken struct {
	UserId    user.Id `db:"user_id"`
	Access    *AccessToken
	Refresh   RefreshToken `db:"refresh_token"`
	ExpiresAt *time.Time   `db:"expires_at"`
}

func NewUserToken(userId user.Id) (*UserToken, error) {
	now := helpers.NowTime()
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
		Access:    &access,
		Refresh:   refresh,
		ExpiresAt: &expiresAt,
	}, nil
}
