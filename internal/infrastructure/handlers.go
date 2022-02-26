package infrastructure

import (
	"bigfood/internal/authorization/actions/auth"
	"bigfood/internal/authorization/actions/refreshToken"
	"bigfood/internal/authorization/actions/sendSmsCode"
	"bigfood/internal/authorization/actions/userLogout"
	"bigfood/internal/cafe/actions/create"
	cafeUserAdd "bigfood/internal/cafeUser/actions/create"
	cafeUserDelete "bigfood/internal/cafeUser/actions/delete"
	cafeUserEdit "bigfood/internal/cafeUser/actions/edit"
	cafeUserList "bigfood/internal/cafeUser/actions/list"
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
	CafeUserListHandler    *cafeUserList.Handler
	CafeUserDeleteHandler  *cafeUserDelete.Handler
	CafeUserEditHandler    *cafeUserEdit.Handler
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
			repositories.PermissionRepository,
			services.UserService,
		),
		RefreshTokenHandler: refreshToken.New(
			repositories.UserTokenRepository,
			repositories.PermissionRepository,
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
		CafeUserListHandler: cafeUserList.New(
			repositories.UserRepository,
			repositories.CafeUserRepository,
		),
		CafeUserDeleteHandler: cafeUserDelete.New(
			repositories.CafeUserRepository,
		),
		CafeUserEditHandler: cafeUserEdit.New(
			repositories.UserRepository,
			repositories.CafeUserRepository,
		),
	}
}
