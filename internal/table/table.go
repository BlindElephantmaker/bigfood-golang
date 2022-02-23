package table

import (
	"bigfood/internal/helpers"
	"errors"
)

type Table struct {
	Id      helpers.Uuid `json:"id" example:"uuid" db:"id"`
	CafeId  helpers.Uuid `json:"cafe-id" example:"uuid" db:"cafe_id"`
	Title   Title        `json:"title" example:"serial number" db:"title"`
	Comment Comment      `json:"comment" example:"comment" db:"comment"`
	Seats   int          `json:"seats" example:"4" db:"seats"`
}

func NewTable(cafeId helpers.Uuid, title Title) *Table {
	comment, _ := NewComment("")
	return &Table{
		Id:      helpers.UuidGenerate(),
		CafeId:  cafeId,
		Title:   title,
		Comment: comment,
		Seats:   4,
	}
}

type Title string
type Comment string

func NewTitle(title string) (Title, error) {
	if len(title) > 32 {
		return "", errors.New("table title is too long") // todo to var
	}

	return Title(title), nil
}

func NewComment(comment string) (Comment, error) {
	if len(comment) > 32 {
		return "", errors.New("table comment is too long") // todo to var
	}

	return Comment(comment), nil
}
