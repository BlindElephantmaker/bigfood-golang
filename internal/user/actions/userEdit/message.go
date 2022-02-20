package userEdit

import "github.com/google/uuid"

type Message struct {
	Id   *uuid.UUID `swaggerignore:"true"`
	Name string     `json:"name" binding:"required" example:"New user name"`
}
