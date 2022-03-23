package auth

import (
	"bigfood/internal/authorization/smsCode"
	"bigfood/internal/cafeUser/permissions"
	"bigfood/internal/user"
	"bigfood/internal/user/userToken"
	"errors"
)

type Message struct {
	Phone   user.Phone      `json:"phone" binding:"required" example:"+71234567890"`
	SmsCode smsCode.SmsCode `json:"sms-code" binding:"required" example:"1234"`
}

type Response struct {
	UserToken *userToken.UserToken
	IsNew     bool
}

var ErrorSmsCodeNotConfirmed = errors.New("sms code not confirmed")

func (h *Handler) Run(m *Message) (*Response, error) {
	err := h.validateSmsCode(m)
	if err != nil {
		return nil, err
	}

	u, err := h.userService.GetOrNewUser(m.Phone)
	if err != nil {
		return nil, err
	}

	token, err := h.createToken(u.Id)
	if err != nil {
		return nil, err
	}

	return &Response{
		UserToken: token,
		IsNew:     u.IsNew(),
	}, nil
}

func (h *Handler) validateSmsCode(m *Message) error {
	confirmCode, err := h.smsCodeRepository.Get(m.Phone)
	if err != nil {
		return err
	}
	if !confirmCode.Compare(m.SmsCode) {
		return ErrorSmsCodeNotConfirmed
	}

	return nil
}

func (h *Handler) createToken(userId user.Id) (*userToken.UserToken, error) {
	userPermissions, err := h.permissionsRepository.GetPermissions(userId)
	if err != nil {
		return nil, err
	}

	token, err := userToken.NewUserToken(userPermissions)
	if err != nil {
		return nil, err
	}

	err = h.tokenRepository.Add(token)
	if err != nil {
		return nil, err
	}

	return token, nil
}

type Handler struct {
	smsCodeRepository     smsCode.Repository
	userRepository        user.Repository
	tokenRepository       userToken.Repository
	permissionsRepository permissions.Repository
	userService           *user.Service
}

func New(
	smsCodeRepository smsCode.Repository,
	users user.Repository,
	tokens userToken.Repository,
	permissions permissions.Repository,
	userService *user.Service,
) *Handler {
	return &Handler{
		smsCodeRepository:     smsCodeRepository,
		userRepository:        users,
		tokenRepository:       tokens,
		permissionsRepository: permissions,
		userService:           userService,
	}
}
