package table

import (
	"bigfood/internal/helpers"
	"errors"
)

type Table struct {
	Id      helpers.Uuid
	CafeId  helpers.Uuid
	Title   Title
	Comment Comment
	Seats   int
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
