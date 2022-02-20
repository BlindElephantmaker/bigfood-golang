package infrastructure

import (
	"bigfood/internal/authorization/smsCode"
	"bigfood/internal/authorization/userToken"
	"bigfood/internal/organization"
	"bigfood/internal/user"
	"github.com/jmoiron/sqlx"
)

type Repositories struct {
	SmsCodeRepository      smsCode.Repository
	UserRepository         user.Repository
	UserTokenRepository    userToken.Repository
	OrganizationRepository organization.Repository
}

func NewRepositories(db *sqlx.DB) *Repositories {
	return &Repositories{
		SmsCodeRepository:      smsCode.NewRepositoryDummy(),
		UserRepository:         user.NewRepositoryPSQL(db),
		UserTokenRepository:    userToken.NewRepositoryPSQL(db),
		OrganizationRepository: organization.NewRepositoryPSQL(db),
	}
}
