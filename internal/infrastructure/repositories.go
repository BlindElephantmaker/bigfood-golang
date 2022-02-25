package infrastructure

import (
	"bigfood/internal/authorization/smsCode"
	"bigfood/internal/cafe"
	"bigfood/internal/cafeUser"
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
}

func NewRepositories(db *sqlx.DB) *Repositories {
	cafeUserRepository := cafeUser.NewRepositoryPSQL(db)

	return &Repositories{
		SmsCodeRepository:   smsCode.NewRepositoryDummy(),
		UserRepository:      user.NewRepositoryPSQL(db),
		UserTokenRepository: userToken.NewRepositoryPSQL(db),
		CafeRepository:      cafe.NewRepositoryPSQL(db, cafeUserRepository),
		CafeUserRepository:  cafeUserRepository,
		TableRepository:     table.NewRepositoryPSQL(db),
	}
}
