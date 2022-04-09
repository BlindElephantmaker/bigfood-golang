package infrastructure

import (
	"bigfood/internal/authorization/action/auth"
	"bigfood/internal/authorization/action/refreshToken"
	"bigfood/internal/authorization/action/sendSmsCode"
	"bigfood/internal/cafe/actions/create"
	"bigfood/internal/cafeUser/actions/create"
	"bigfood/internal/cafeUser/actions/delete"
	"bigfood/internal/cafeUser/actions/edit"
	"bigfood/internal/cafeUser/actions/list"
	"bigfood/internal/reserve/actions/create"
	"bigfood/internal/reserve/actions/delete"
	"bigfood/internal/reserve/actions/edit"
	"bigfood/internal/reserve/actions/get"
	"bigfood/internal/reserve/actions/historyByTable"
	"bigfood/internal/reserve/actions/listByTable"
	"bigfood/internal/reserve/actions/undelete"
	"bigfood/internal/table/actions/createMass"
	"bigfood/internal/table/actions/getList"
	"bigfood/internal/table/actions/tableCreate"
	"bigfood/internal/table/actions/tableDelete"
	"bigfood/internal/table/actions/tableDeleteAll"
	"bigfood/internal/table/actions/tableEdit"
	"bigfood/internal/user/actions/userEdit"
)

type Handlers struct {
	SendSmsCode                  *sendSmsCode.Handler
	UserAuthHandler              *auth.Handler
	RefreshTokenHandler          *refreshToken.Handler
	UserEditHandler              *userEdit.Handler
	CafeCreateHandler            *createCafe.Handler
	TableCreateHandler           *tableCreate.Handler
	TableCreateMassHandler       *createMass.Handler
	TableGetListHandler          *getList.Handler
	TableEditHandler             *tableEdit.Handler
	TableDeleteHandler           *tableDelete.Handler
	TableDeleteAllHandler        *tableDeleteAll.Handler
	CafeUserCreateHandler        *cafeUserCreate.Handler
	CafeUserListHandler          *cafeUserList.Handler
	CafeUserDeleteHandler        *cafeUserDelete.Handler
	CafeUserEditHandler          *cafeUserEdit.Handler
	ReserveCreateHandler         *reserveCreate.Handler
	ReserveGetHandler            *reserveGet.Handler
	ReserveDeleteHandler         *reserveDelete.Handler
	ReserveUndeleteHandler       *reserveUndelete.Handler
	ReserveEditHandler           *reserveEdit.Handler
	ReserveListByTableHandler    *reserveListByTable.Handler
	ReserveHistoryByTableHandler *reserveHistoryByTable.Handler
}

func NewHandlers(repositories *Repositories, services *Services) *Handlers {
	return &Handlers{
		SendSmsCode: sendSmsCode.New(
			services.SmsCodeService,
			repositories.SmsCodeRepository,
		),
		UserAuthHandler: auth.New(
			repositories.SmsCodeRepository,
			services.UserService,
			services.UserTokenService,
		),
		RefreshTokenHandler: refreshToken.New(
			repositories.UserTokenRepository,
			services.UserTokenService,
		),
		UserEditHandler: userEdit.New(
			repositories.UserRepository,
		),
		CafeCreateHandler: createCafe.New(
			repositories.CafeRepository,
			repositories.CafeUserRepository,
			services.Transactions,
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
		CafeUserCreateHandler: cafeUserCreate.New(
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
		ReserveCreateHandler: reserveCreate.New(
			repositories.ReserveRepository,
			services.reserveActionHelper,
		),
		ReserveGetHandler: reserveGet.New(
			repositories.ReserveRepository,
		),
		ReserveDeleteHandler: reserveDelete.New(
			repositories.ReserveRepository,
		),
		ReserveUndeleteHandler: reserveUndelete.New(
			repositories.ReserveRepository,
		),
		ReserveEditHandler: reserveEdit.New(
			repositories.ReserveRepository,
			services.reserveActionHelper,
		),
		ReserveListByTableHandler: reserveListByTable.New(
			repositories.ReserveRepository,
		),
		ReserveHistoryByTableHandler: reserveHistoryByTable.New(
			repositories.ReserveRepository,
		),
	}
}
