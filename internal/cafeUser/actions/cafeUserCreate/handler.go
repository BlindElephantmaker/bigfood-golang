package cafeUserCreate

import (
	"bigfood/internal/cafe"
	"bigfood/internal/cafeUser"
	"bigfood/internal/cafeUser/actions"
	"bigfood/internal/helpers"
	"bigfood/internal/user"
)

type Message struct {
	CafeId  cafe.Id           `json:"cafe-id" binding:"required" example:"uuid"`
	Phone   helpers.Phone     `json:"phone" binding:"required" example:"User phone"`
	Comment *cafeUser.Comment `json:"comment"`
	Roles   []string          `json:"roles" binding:"required" example:"owner,admin,hostess"` // todo: collection
}

var errorCafeUserAlreadyExist = helpers.ErrorUnprocessableEntity("cafe user already exist")

func (h *Handler) Run(m *Message) (*actions.Response, error) {
	comment, roles, err := parseMessage(m)
	if err != nil {
		return nil, err
	}

	usr, err := h.UserService.GetOrNewUser(m.Phone)
	if err != nil {
		return nil, err
	}

	cafeUsr, err := h.CafeUserRepository.GetByCafeAndUser(m.CafeId, usr.Id)
	if err == cafeUser.ErrorNoResult {
		cafeUsr = cafeUser.NewCafeUser(m.CafeId, usr.Id, comment)
		err = h.CafeUserRepository.Add(cafeUsr, helpers.NowTime(), roles)
	} else if err != nil {
		return nil, err
	} else {
		if !cafeUsr.IsDeleted() {
			return nil, errorCafeUserAlreadyExist
		}
		cafeUsr.DeletedAt = nil
		cafeUsr.Comment = comment
		err = h.CafeUserRepository.Update(cafeUsr, roles)
	}

	return &actions.Response{
		CafeUser: cafeUsr,
		Roles:    roles,
		UserName: usr.Name,
	}, err
}

func parseMessage(m *Message) (cafeUser.Comment, cafeUser.Roles, error) {
	roles, err := cafeUser.ParseRoles(m.Roles)
	if err != nil {
		return "", nil, err
	}

	var comment cafeUser.Comment
	if m.Comment != nil { // todo: maybe move it to UnmarshalJSON?
		comment = *m.Comment
	} else {
		comment = cafeUser.NewComment()
	}

	return comment, roles, nil
}

type Handler struct {
	CafeUserRepository cafeUser.Repository
	UserService        *user.Service
}

func New(cafeUsers cafeUser.Repository, userService *user.Service) *Handler {
	return &Handler{
		CafeUserRepository: cafeUsers,
		UserService:        userService,
	}
}
