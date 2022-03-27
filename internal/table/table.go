package table

import (
	"bigfood/internal/cafe"
)

type Table struct {
	Id      Id      `json:"id" example:"uuid" db:"id"`
	CafeId  cafe.Id `json:"cafe-id" example:"uuid" db:"cafe_id"`
	Title   Title   `json:"title" example:"serial number" db:"title"`
	Comment Comment `json:"comment" example:"comment" db:"comment"`
	Seats   Seats   `json:"seats" example:"4" db:"seats"`
}

func NewTable(cafeId cafe.Id, title Title) *Table {
	return &Table{
		Id:      NewId(),
		CafeId:  cafeId,
		Title:   title,
		Comment: NewComment(),
		Seats:   NewSeats(),
	}
}
