package auth

import (
	"bigfood/internal/authorization/smsCode"
	"bigfood/internal/helpers"
	"bigfood/internal/user"
	"bigfood/internal/user/userToken"
)

type Message struct {
	Phone   helpers.Phone   `json:"phone" binding:"required" example:"+71234567890"`
	SmsCode smsCode.SmsCode `json:"sms-code" binding:"required" example:"1234"`
}

type Response struct {
	IsNew   bool                   `json:"is-new"`
	UserId  user.Id                `json:"user-id" example:"UUID"`
	Access  userToken.AccessToken  `json:"access-token"`
	Refresh userToken.RefreshToken `json:"refresh-token" example:"UUID"`
}

var errorSmsCodeNotConfirmed = helpers.ErrorUnprocessableEntity("sms code not confirmed")

func (h *Handler) Run(m *Message) (*Response, error) {
	err := h.validateSmsCode(m)
	if err != nil {
		return nil, err
	}

	usr, err := h.userService.GetOrNewUser(m.Phone)
	if err != nil {
		return nil, err
	}

	access, usrToken, err := h.userTokenService.CreateTokens(usr.Id)
	if err != nil {
		return nil, err
	}

	return &Response{
		UserId:  usr.Id,
		IsNew:   usr.IsNew(),
		Refresh: usrToken.Refresh,
		Access:  access,
	}, nil
}

func (h *Handler) validateSmsCode(m *Message) error {
	confirmCode, err := h.smsCodeRepository.Get(m.Phone)
	if err != nil {
		return err
	}
	if !confirmCode.Compare(m.SmsCode) {
		return errorSmsCodeNotConfirmed
	}

	return nil
}

type Handler struct {
	smsCodeRepository smsCode.Repository
	userService       *user.Service
	userTokenService  *userToken.Service
}

func New(
	smsCodeRepository smsCode.Repository,
	userService *user.Service,
	userTokenService *userToken.Service,
) *Handler {
	return &Handler{
		smsCodeRepository: smsCodeRepository,
		userService:       userService,
		userTokenService:  userTokenService,
	}
}
