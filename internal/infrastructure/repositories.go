package infrastructure

import (
	"bigfood/internal/authorization/smsCode"
	"bigfood/internal/cafe"
	"bigfood/internal/cafeUser"
	"bigfood/internal/contact"
	"bigfood/internal/reserve"
	"bigfood/internal/table"
	"bigfood/internal/user"
	"bigfood/internal/user/userToken"
	"github.com/jmoiron/sqlx"
)

type Repositories struct {
	SmsCodeRepository   smsCode.Repository
	UserRepository      user.Repository
	UserTokenRepository userToken.Repository
	CafeRepository      cafe.Repository
	CafeUserRepository  cafeUser.Repository
	TableRepository     table.Repository
	ContactRepository   contact.Repository
	ReserveRepository   reserve.Repository
}

func NewRepositories(db *sqlx.DB) *Repositories {
	return &Repositories{
		SmsCodeRepository:   smsCode.NewRepositoryDummy(),
		UserRepository:      user.NewRepositoryPsql(db),
		UserTokenRepository: userToken.NewRepositoryPsql(db),
		CafeRepository:      cafe.NewRepositoryPsql(db),
		CafeUserRepository:  cafeUser.NewRepositoryPsql(db),
		TableRepository:     table.NewRepositoryPsql(db),
		ContactRepository:   contact.NewRepositoryPsql(db),
		ReserveRepository:   reserve.NewRepositoryPsql(db),
	}
}
