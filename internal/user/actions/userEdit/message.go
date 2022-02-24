package userEdit

import (
	"bigfood/internal/helpers"
)

type Message struct {
	Id   helpers.Uuid `swaggerignore:"true"`
	Name string       `json:"name" binding:"required" example:"New user name"`
	// todo: edit photo
}
