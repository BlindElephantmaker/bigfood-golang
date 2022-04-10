package infrastructure

import (
	"bigfood/internal/authorization/smsCode"
	"bigfood/internal/cafe"
	"bigfood/internal/cafeUser"
	"bigfood/internal/contact"
	"bigfood/internal/reserve"
	"bigfood/internal/table"
	"bigfood/internal/table/actions/tableListAvailable"
	"bigfood/internal/user"
	"bigfood/internal/user/userToken"
	"github.com/jmoiron/sqlx"
)

type Repositories struct {
	SmsCode            smsCode.Repository
	User               user.Repository
	UserToken          userToken.Repository
	Cafe               cafe.Repository
	CafeUser           cafeUser.Repository
	Table              table.Repository
	TableListAvailable tableListAvailable.Repository
	Contact            contact.Repository
	Reserve            reserve.Repository
}

func NewRepositories(db *sqlx.DB) *Repositories {
	return &Repositories{
		SmsCode:            smsCode.NewRepositoryDummy(),
		User:               user.NewRepositoryPsql(db),
		UserToken:          userToken.NewRepositoryPsql(db),
		Cafe:               cafe.NewRepositoryPsql(db),
		CafeUser:           cafeUser.NewRepositoryPsql(db),
		Table:              table.NewRepositoryPsql(db),
		TableListAvailable: tableListAvailable.NewRepositoryPsql(db),
		Contact:            contact.NewRepositoryPsql(db),
		Reserve:            reserve.NewRepositoryPsql(db),
	}
}
