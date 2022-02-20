package userLogout

import "github.com/google/uuid"

type Message struct {
	UserId *uuid.UUID `swaggerignore:"true"`
	Token  string     `json:"token" binding:"required" example:"uuid"`
}
