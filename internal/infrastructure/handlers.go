package infrastructure

import (
	"bigfood/internal/authorization/action/auth"
	"bigfood/internal/authorization/action/refreshToken"
	"bigfood/internal/authorization/action/sendSmsCode"
	"bigfood/internal/cafe/actions/create"
	"bigfood/internal/cafeUser/actions/cafeUserList"
	"bigfood/internal/cafeUser/actions/create"
	"bigfood/internal/cafeUser/actions/delete"
	"bigfood/internal/cafeUser/actions/edit"
	"bigfood/internal/reserve/actions/create"
	"bigfood/internal/reserve/actions/delete"
	"bigfood/internal/reserve/actions/edit"
	"bigfood/internal/reserve/actions/get"
	"bigfood/internal/reserve/actions/historyByTable"
	"bigfood/internal/reserve/actions/listByTable"
	"bigfood/internal/reserve/actions/undelete"
	"bigfood/internal/table/actions/createMass"
	"bigfood/internal/table/actions/tableCreate"
	"bigfood/internal/table/actions/tableDelete"
	"bigfood/internal/table/actions/tableDeleteAll"
	"bigfood/internal/table/actions/tableEdit"
	"bigfood/internal/table/actions/tableList"
	"bigfood/internal/table/actions/tableListAvailable"
	"bigfood/internal/user/actions/userEdit"
)

type Handlers struct {
	SendSmsCode           *sendSmsCode.Handler
	UserAuth              *auth.Handler
	RefreshToken          *refreshToken.Handler
	UserEdit              *userEdit.Handler
	CafeCreate            *createCafe.Handler
	TableCreate           *tableCreate.Handler
	TableCreateMass       *createMass.Handler
	TableList             *tableList.Handler
	TableListAvailable    *tableListAvailable.Handler
	TableEdit             *tableEdit.Handler
	TableDelete           *tableDelete.Handler
	TableDeleteAll        *tableDeleteAll.Handler
	CafeUserCreate        *cafeUserCreate.Handler
	CafeUserList          *cafeUserList.Handler
	CafeUserDelete        *cafeUserDelete.Handler
	CafeUserEdit          *cafeUserEdit.Handler
	ReserveCreate         *reserveCreate.Handler
	ReserveGet            *reserveGet.Handler
	ReserveDelete         *reserveDelete.Handler
	ReserveUndelete       *reserveUndelete.Handler
	ReserveEdit           *reserveEdit.Handler
	ReserveListByTable    *reserveListByTable.Handler
	ReserveHistoryByTable *reserveHistoryByTable.Handler
}

func NewHandlers(repositories *Repositories, services *Services) *Handlers {
	return &Handlers{
		SendSmsCode: sendSmsCode.New(
			services.SmsCode,
			repositories.SmsCode,
		),
		UserAuth: auth.New(
			repositories.SmsCode,
			services.User,
			services.UserToken,
		),
		RefreshToken: refreshToken.New(
			repositories.UserToken,
			services.UserToken,
		),
		UserEdit: userEdit.New(
			repositories.User,
		),
		CafeCreate: createCafe.New(
			repositories.Cafe,
			repositories.CafeUser,
			services.Transactions,
		),
		TableCreateMass: createMass.New(
			repositories.Table,
		),
		TableCreate: tableCreate.New(
			repositories.Table,
		),
		TableList: tableList.New(
			repositories.Table,
		),
		TableEdit: tableEdit.New(
			repositories.Table,
		),
		TableDelete: tableDelete.New(
			repositories.Table,
		),
		TableDeleteAll: tableDeleteAll.New(
			repositories.Table,
		),
		CafeUserCreate: cafeUserCreate.New(
			repositories.CafeUser,
			services.User,
		),
		CafeUserList: cafeUserList.New(
			repositories.User,
			repositories.CafeUser,
		),
		CafeUserDelete: cafeUserDelete.New(
			repositories.CafeUser,
		),
		CafeUserEdit: cafeUserEdit.New(
			repositories.User,
			repositories.CafeUser,
		),
		ReserveCreate: reserveCreate.New(
			repositories.Reserve,
			services.reserveActionHelper,
		),
		ReserveGet: reserveGet.New(
			repositories.Reserve,
		),
		ReserveDelete: reserveDelete.New(
			repositories.Reserve,
		),
		ReserveUndelete: reserveUndelete.New(
			repositories.Reserve,
		),
		ReserveEdit: reserveEdit.New(
			repositories.Reserve,
			services.reserveActionHelper,
		),
		ReserveListByTable: reserveListByTable.New(
			repositories.Reserve,
		),
		ReserveHistoryByTable: reserveHistoryByTable.New(
			repositories.Reserve,
		),
		TableListAvailable: tableListAvailable.New(
			repositories.TableListAvailable,
		),
	}
}
