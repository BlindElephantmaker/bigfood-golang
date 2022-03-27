package cafeUser

import (
	"bigfood/internal/cafe"
	"bigfood/internal/user"
	"time"
)

type CafeUser struct {
	Id        Id         `json:"id" example:"uuid" db:"id"`
	CafeId    cafe.Id    `json:"cafe-id" example:"uuid" db:"cafe_id"`
	UserId    user.Id    `json:"user-id" example:"uuid" db:"user_id"`
	Comment   Comment    `json:"comment" db:"comment"`
	DeletedAt *time.Time `json:"-" db:"deleted_at" swaggerignore:"true"`
}

func NewCafeUser(cafeId cafe.Id, userId user.Id, comment Comment) *CafeUser {
	return &CafeUser{
		Id:        NewId(),
		CafeId:    cafeId,
		UserId:    userId,
		Comment:   comment,
		DeletedAt: nil,
	}
}

func (cu *CafeUser) IsDeleted() bool {
	return cu.DeletedAt != nil
}
