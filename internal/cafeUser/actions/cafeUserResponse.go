package actions

import (
	"bigfood/internal/cafeUser"
	"bigfood/internal/user"
)

type Response struct {
	*cafeUser.CafeUser
	cafeUser.Roles `json:"roles" example:"owner,admin,hostess"` // todo: strange in swagger
	user.Name      `json:"name"`
}
