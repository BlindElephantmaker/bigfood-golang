package cafeUser

import (
	"bigfood/internal/cafeUser/role"
	"bigfood/internal/helpers"
	"time"
)

type User struct {
	Id        helpers.Uuid `json:"id" example:"uuid" db:"id"`
	CafeId    helpers.Uuid `json:"cafe-id" example:"uuid" db:"cafe_id"`
	UserId    helpers.Uuid `json:"user-id" example:"uuid" db:"user_id"`
	Comment   Comment      `json:"comment" db:"comment"`
	Roles     []role.Role  `json:"roles" example:"owner,admin,hostess"`
	DeletedAt *time.Time   `json:"-" db:"deleted_at" swaggerignore:"true"`
}

func NewCafeUser(cafeId, userId helpers.Uuid, comment Comment, roles role.Roles) *User {
	return &User{
		Id:        helpers.UuidGenerate(),
		CafeId:    cafeId,
		UserId:    userId,
		Comment:   comment,
		Roles:     roles,
		DeletedAt: nil,
	}
}

func CreateCafeUser(id, cafeId, userId helpers.Uuid, comment Comment, deletedAt *time.Time, roles role.Roles) *User {
	return &User{
		Id:        id,
		CafeId:    cafeId,
		UserId:    userId,
		Comment:   comment,
		Roles:     roles,
		DeletedAt: deletedAt,
	}
}

func (cu *User) IsDeleted() bool {
	return cu.DeletedAt != nil
}
