package infrastructure

import "bigfood/internal/authorization/smsCode"

type Services struct {
	SmsCodeService smsCode.Service
}

func NewServices() *Services {
	return &Services{
		SmsCodeService: smsCode.NewServiceDummy(),
	}
}
