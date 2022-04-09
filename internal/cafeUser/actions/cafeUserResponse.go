package actions

import (
	"bigfood/internal/cafeUser"
	"bigfood/internal/user"
)

type Response struct {
	CafeUser *cafeUser.CafeUser
	Roles    cafeUser.Roles `json:"roles" example:"owner,admin,hostess"` // todo: get list or roles from const
	UserName user.Name      `json:"user-name"`
}
