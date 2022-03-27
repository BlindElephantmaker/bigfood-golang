package infrastructure

import (
	"bigfood/internal/authorization/smsCode"
	"bigfood/internal/user"
	"bigfood/pkg/database"
	"github.com/jmoiron/sqlx"
)

type Services struct {
	SmsCodeService smsCode.Service
	UserService    *user.Service
	Transactions   *database.TransactionFactory
}

func NewServices(repositories *Repositories, db *sqlx.DB) *Services {
	return &Services{
		SmsCodeService: smsCode.NewServiceDummy(),
		UserService:    user.NewService(repositories.UserRepository),
		Transactions:   database.NewTransactionFactory(db),
	}
}
