package userLogout

import (
	"bigfood/internal/helpers"
)

type Message struct {
	UserId helpers.Uuid `swaggerignore:"true"`
	Token  string       `json:"token" binding:"required" example:"uuid"`
}
