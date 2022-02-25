package infrastructure

import (
	"bigfood/internal/authorization/actions/auth"
	"bigfood/internal/authorization/actions/refreshToken"
	"bigfood/internal/authorization/actions/sendSmsCode"
	"bigfood/internal/authorization/actions/userLogout"
	"bigfood/internal/cafe/actions/create"
	cafeUserAdd "bigfood/internal/cafeUser/actions/create"
	"bigfood/internal/table/actions/createMass"
	"bigfood/internal/table/actions/getList"
	"bigfood/internal/table/actions/tableCreate"
	"bigfood/internal/table/actions/tableDelete"
	"bigfood/internal/table/actions/tableDeleteAll"
	"bigfood/internal/table/actions/tableEdit"
	"bigfood/internal/user/actions/userEdit"
)

type Handlers struct {
	SendSmsCode            *sendSmsCode.Handler
	UserAuthHandler        *auth.Handler
	RefreshTokenHandler    *refreshToken.Handler
	UserLogoutHandler      *userLogout.Handler
	UserEditHandler        *userEdit.Handler
	CafeCreateHandler      *createCafe.Handler
	TableCreateHandler     *tableCreate.Handler
	TableCreateMassHandler *createMass.Handler
	TableGetListHandler    *getList.Handler
	TableEditHandler       *tableEdit.Handler
	TableDeleteHandler     *tableDelete.Handler
	TableDeleteAllHandler  *tableDeleteAll.Handler
	CafeUserAddHandler     *cafeUserAdd.Handler
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
			services.UserService,
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
		TableCreateHandler: tableCreate.New(
			repositories.TableRepository,
		),
		TableGetListHandler: getList.New(
			repositories.TableRepository,
		),
		TableEditHandler: tableEdit.New(
			repositories.TableRepository,
		),
		TableDeleteHandler: tableDelete.New(
			repositories.TableRepository,
		),
		TableDeleteAllHandler: tableDeleteAll.New(
			repositories.TableRepository,
		),
		CafeUserAddHandler: cafeUserAdd.New(
			repositories.CafeUserRepository,
			services.UserService,
		),
	}
}
