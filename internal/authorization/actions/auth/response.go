package auth

import (
	"bigfood/internal/user/userToken"
)

type Response struct {
	UserToken *userToken.UserToken
	IsNew     bool
}
