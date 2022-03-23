package sendSmsCode

import (
	"bigfood/internal/authorization/smsCode"
	"bigfood/internal/user"
	"errors"
	"time"
)

type Message struct {
	Phone user.Phone `json:"phone" binding:"required" example:"+71234567890"`
	// todo: captcha
}

const (
	maxRetryCount = 3
	ttl           = time.Minute * 30
)

var ErrorRetryCountExceeded = errors.New("retry count of sms code requests exceeded")

type Handler struct {
	smsCodeService    smsCode.Service
	smsCodeRepository smsCode.Repository
}

func New(service smsCode.Service, repository smsCode.Repository) *Handler {
	return &Handler{
		smsCodeService:    service,
		smsCodeRepository: repository,
	}
}

func (h *Handler) Run(m Message) error {
	count, err := h.smsCodeRepository.Count(m.Phone)
	if err != nil {
		return err
	}
	if count >= maxRetryCount {
		return ErrorRetryCountExceeded
	}

	code := smsCode.New()
	err = h.smsCodeRepository.Add(code, m.Phone, ttl)
	if err != nil {
		return err
	}

	err = h.smsCodeService.Send(string(code), m.Phone)
	if err != nil {
		_ = h.smsCodeRepository.DeleteLast(m.Phone)
		return err
	}

	return nil
}
