package userToken

import (
	"bigfood/internal/helpers"
	"bigfood/internal/user"
	"time"
)

const refreshTTL = helpers.Year

type UserToken struct {
	UserId    user.Id      `db:"user_id"`
	Refresh   RefreshToken `db:"refresh_token"`
	ExpiresAt time.Time    `db:"expires_at"`
}

func newUserToken(userId user.Id, now time.Time) *UserToken {
	return &UserToken{
		UserId:    userId,
		Refresh:   newRefresh(),
		ExpiresAt: now.Add(refreshTTL),
	}
}
