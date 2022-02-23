package table

import "bigfood/internal/helpers"

type Table struct {
	Id      helpers.Uuid `json:"id" example:"uuid" db:"id"`
	CafeId  helpers.Uuid `json:"cafe-id" example:"uuid" db:"cafe_id"`
	Title   Title        `json:"title" example:"serial number" db:"title"`
	Comment Comment      `json:"comment" example:"comment" db:"comment"`
	Seats   Seats        `json:"seats" example:"4" db:"seats"`
}

func NewTable(cafeId helpers.Uuid, title Title) *Table {
	comment, _ := ParseComment("")
	return &Table{
		Id:      helpers.UuidGenerate(),
		CafeId:  cafeId,
		Title:   title,
		Comment: comment,
		Seats:   NewSeats(),
	}
}
