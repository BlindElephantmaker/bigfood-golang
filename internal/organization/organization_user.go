package organization

import (
	"bigfood/internal/helpers/timeHelper"
	"github.com/google/uuid"
	"time"
)

type User struct {
	Id             *uuid.UUID
	OrganizationId *uuid.UUID
	UserId         *uuid.UUID
	CreatedAt      *time.Time
	Roles          []Role
}

func NewOrganizationUser(organizationId, userId *uuid.UUID) *User {
	id := uuid.New()
	now := timeHelper.Now()
	return &User{
		Id:             &id,
		OrganizationId: organizationId,
		UserId:         userId,
		CreatedAt:      &now,
		Roles: []Role{
			RoleOwner,
			RoleAdmin,
			RoleHostess,
		},
	}
}
