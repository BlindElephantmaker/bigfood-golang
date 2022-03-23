package userEdit

import (
	"bigfood/internal/helpers"
	"bigfood/internal/user"
)

type Message struct {
	Id   helpers.Uuid `swaggerignore:"true"`
	Name user.Name    `json:"name" binding:"required" example:"New user name"`
	// todo: edit photo
}

func (h *Handler) Run(message *Message) error {
	u, err := h.userRepository.Get(message.Id)
	if err != nil {
		return err
	}
	u.Name = message.Name
	return h.userRepository.Update(u)
}

type Handler struct {
	userRepository user.Repository
}

func New(users user.Repository) *Handler {
	return &Handler{
		userRepository: users,
	}
}
