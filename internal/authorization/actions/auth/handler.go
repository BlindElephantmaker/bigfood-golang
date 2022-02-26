package auth

import (
	"bigfood/internal/authorization/smsCode"
	"bigfood/internal/cafeUser/permissions"
	"bigfood/internal/helpers"
	"bigfood/internal/user"
	"bigfood/internal/user/userToken"
	"errors"
)

type Handler struct {
	smsCodeRepository     smsCode.Repository
	userRepository        user.Repository
	tokenRepository       userToken.Repository
	permissionsRepository permissions.Repository
	userService           *user.Service
}

var ErrorSmsCodeNotConfirmed = errors.New("sms code not confirmed")

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

func (h *Handler) Run(message *Message) (*Response, error) {
	phone, err := user.NewPhone(message.Phone)
	if err != nil {
		return nil, err
	}
	code, err := smsCode.Parse(message.SmsCode)
	if err != nil {
		return nil, err
	}

	err = h.validateSmsCode(phone, code)
	if err != nil {
		return nil, err
	}

	u, err := h.userService.GetOrNewUser(phone)
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

func (h *Handler) validateSmsCode(phone user.Phone, code smsCode.Code) error {
	confirmCode, err := h.smsCodeRepository.Get(phone)
	if err != nil {
		return err
	}
	if !confirmCode.Compare(code) {
		return ErrorSmsCodeNotConfirmed
	}

	return nil
}

func (h *Handler) createToken(id helpers.Uuid) (*userToken.UserToken, error) {
	userPermissions, err := h.permissionsRepository.GetPermissions(id)
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
