package auth

import "bigfood/internal/authorization/userToken"

type Response struct {
	UserToken *userToken.UserToken
	IsNew     bool
}
