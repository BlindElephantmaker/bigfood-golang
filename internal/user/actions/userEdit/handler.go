package userEdit

import (
	"bigfood/internal/user"
	"errors"
)

var ErrorUserNameIsEmpty = errors.New("username is empty")

type Handler struct {
	userRepository user.Repository
}

func New(users user.Repository) *Handler {
	return &Handler{
		userRepository: users,
	}
}

func (h *Handler) Run(message *Message) error {
	if message.Name == "" {
		return ErrorUserNameIsEmpty
	}

	name, err := user.NewName(message.Name)
	if err != nil {
		return err
	}

	u, err := h.userRepository.Get(message.Id)
	if err != nil {
		return err
	}

	u.Name = name

	return h.userRepository.Update(u)
}
