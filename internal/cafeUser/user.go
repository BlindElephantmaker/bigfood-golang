package cafeUser

import (
	"bigfood/internal/helpers"
	"time"
)

type CafeUser struct {
	Id        helpers.Uuid `json:"id" example:"uuid" db:"id"`
	CafeId    helpers.Uuid `json:"cafe-id" example:"uuid" db:"cafe_id"`
	UserId    helpers.Uuid `json:"user-id" example:"uuid" db:"user_id"`
	Comment   Comment      `json:"comment" db:"comment"`
	DeletedAt *time.Time   `json:"-" db:"deleted_at" swaggerignore:"true"`
}

func NewCafeUser(cafeId, userId helpers.Uuid, comment Comment) *CafeUser {
	return &CafeUser{
		Id:        helpers.UuidGenerate(),
		CafeId:    cafeId,
		UserId:    userId,
		Comment:   comment,
		DeletedAt: nil,
	}
}

func (cu *CafeUser) IsDeleted() bool {
	return cu.DeletedAt != nil
}
