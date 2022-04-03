package infrastructure

import (
	"bigfood/internal/authorization/smsCode"
	"bigfood/internal/user"
	"bigfood/internal/user/userToken"
	"bigfood/pkg/database"
	"github.com/jmoiron/sqlx"
)

type Services struct {
	SmsCodeService   smsCode.Service
	UserService      *user.Service
	UserTokenService *userToken.Service
	Transactions     *database.TransactionFactory
}

func NewServices(repositories *Repositories, db *sqlx.DB) *Services {
	return &Services{
		SmsCodeService:   smsCode.NewServiceDummy(),
		UserService:      user.NewService(repositories.UserRepository),
		UserTokenService: userToken.NewService(repositories.UserTokenRepository),
		Transactions:     database.NewTransactionFactory(db),
	}
}
