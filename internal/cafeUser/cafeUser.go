package cafeUser

import (
	"bigfood/internal/cafe"
	"bigfood/internal/user"
	"time"
)

type CafeUser struct {
	Id        Id         `json:"id" db:"id" example:"uuid"`
	CafeId    cafe.Id    `json:"cafe-id" db:"cafe_id" example:"uuid"`
	UserId    user.Id    `json:"user-id" db:"user_id" example:"uuid"`
	Comment   Comment    `json:"comment" db:"comment"`
	DeletedAt *time.Time `json:"-" db:"deleted_at" swaggerignore:"true"`
}

func NewCafeUser(cafeId cafe.Id, userId user.Id, comment Comment) *CafeUser {
	return &CafeUser{
		Id:        newId(),
		CafeId:    cafeId,
		UserId:    userId,
		Comment:   comment,
		DeletedAt: nil,
	}
}

func (cu *CafeUser) IsDeleted() bool {
	return cu.DeletedAt != nil
}
