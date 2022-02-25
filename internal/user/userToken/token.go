package userToken

import (
	"bigfood/internal/cafeUser/role"
	"bigfood/internal/helpers"
	"time"
)

const (
	day        = 24 * time.Hour
	accessTTL  = 1 * day
	refreshTTL = 360 * day
)

type UserToken struct {
	UserId    helpers.Uuid `db:"user_id"`
	Access    *AccessToken
	Refresh   RefreshToken `db:"refresh_token"`
	ExpiresAt *time.Time   `db:"expires_at"`
}

func NewUserToken(permissions *role.Permissions) (*UserToken, error) {
	now := helpers.TimeNow()
	expiresAt := now.Add(refreshTTL)

	access, err := NewAccess(permissions, &now, accessTTL)
	if err != nil {
		return nil, err
	}

	refresh := NewRefresh()
	if err != nil {
		return nil, err
	}

	return &UserToken{
		UserId:    permissions.UserId,
		Access:    &access,
		Refresh:   refresh,
		ExpiresAt: &expiresAt,
	}, nil
}
