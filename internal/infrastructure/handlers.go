package infrastructure

import (
	"bigfood/internal/authorization/actions/auth"
	"bigfood/internal/authorization/actions/refreshToken"
	"bigfood/internal/authorization/actions/sendSmsCode"
	"bigfood/internal/authorization/actions/userLogout"
	"bigfood/internal/cafe/actions/create"
	"bigfood/internal/table/createMass"
	"bigfood/internal/table/getList"
	"bigfood/internal/user/actions/userEdit"
)

type Handlers struct {
	SendSmsCode            *sendSmsCode.Handler
	UserAuthHandler        *auth.Handler
	RefreshTokenHandler    *refreshToken.Handler
	UserLogoutHandler      *userLogout.Handler
	UserEditHandler        *userEdit.Handler
	CafeCreateHandler      *createCafe.Handler
	TableCreateMassHandler *createMass.Handler
	TableGetListHandler    *getList.Handler
}

func NewHandlers(repositories *Repositories, services *Services) *Handlers {
	return &Handlers{
		SendSmsCode: sendSmsCode.New(
			services.SmsCodeService,
			repositories.SmsCodeRepository,
		),
		UserAuthHandler: auth.New(
			repositories.SmsCodeRepository,
			repositories.UserRepository,
			repositories.UserTokenRepository,
			repositories.CafeUserRepository,
		),
		RefreshTokenHandler: refreshToken.New(
			repositories.UserTokenRepository,
			repositories.CafeUserRepository,
		),
		UserLogoutHandler: userLogout.New(
			repositories.UserTokenRepository,
		),
		UserEditHandler: userEdit.New(
			repositories.UserRepository,
		),
		CafeCreateHandler: createCafe.New(
			repositories.CafeRepository,
		),
		TableCreateMassHandler: createMass.New(
			repositories.TableRepository,
		),
		TableGetListHandler: getList.New(
			repositories.TableRepository,
		),
	}
}
