package actions

import (
	"bigfood/internal/cafeUser"
	"bigfood/internal/user"
)

type Response struct {
	*cafeUser.CafeUser
	Roles     cafeUser.Roles `json:"roles" example:"owner,admin,hostess"` // todo: get list or roles from const
	user.Name `json:"name"`
}