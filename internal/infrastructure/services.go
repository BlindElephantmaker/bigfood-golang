package infrastructure

import (
	"bigfood/internal/authorization/smsCode"
	"bigfood/internal/user"
)

type Services struct {
	SmsCodeService smsCode.Service
	UserService    *user.Service
}

func NewServices(repositories *Repositories) *Services {
	return &Services{
		SmsCodeService: smsCode.NewServiceDummy(),
		UserService:    user.NewService(repositories.UserRepository),
	}
}
