package infrastructure

import (
	"bigfood/internal/authorization/smsCode"
	"bigfood/internal/reserve/actions"
	"bigfood/internal/user"
	"bigfood/internal/user/userToken"
	"bigfood/pkg/database"
	"github.com/jmoiron/sqlx"
)

type Services struct {
	SmsCode             smsCode.Service
	User                *user.Service
	UserToken           *userToken.Service
	Transactions        *database.TransactionFactory
	reserveActionHelper *reserveAction.Helper
}

func NewServices(repositories *Repositories, db *sqlx.DB) *Services {
	return &Services{
		SmsCode:      smsCode.NewServiceDummy(),
		User:         user.NewService(repositories.User),
		UserToken:    userToken.NewService(repositories.UserToken),
		Transactions: database.NewTransactionFactory(db),
		reserveActionHelper: reserveAction.NewHelper(
			repositories.Contact,
			repositories.Table,
		),
	}
}
