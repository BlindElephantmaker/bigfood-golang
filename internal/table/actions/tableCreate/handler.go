package tableCreate

import (
	"bigfood/internal/helpers"
	"bigfood/internal/table"
)

type Message struct {
	CafeId  helpers.Uuid   `json:"cafe-id" binding:"required" example:"uuid"`
	Title   *table.Title   `json:"title"`
	Comment *table.Comment `json:"comment"`
	Seats   *table.Seats   `json:"seats"`
}

func (h *Handler) Run(m *Message) (*table.Table, error) {
	var title table.Title
	if m.Title != nil {
		title = *m.Title
	} else {
		title, _ = table.ParseTitle("New table") // todo: serial number
	}

	var comment table.Comment
	if m.Comment != nil {
		comment = *m.Comment
	} else {
		comment = table.NewComment()
	}

	var seats table.Seats
	if m.Seats != nil {
		seats = *m.Seats
	} else {
		seats = table.NewSeats()
	}

	newTable := &table.Table{
		Id:      table.NewId(),
		CafeId:  m.CafeId,
		Title:   title,
		Comment: comment,
		Seats:   seats,
	}

	oneTableList := []*table.Table{newTable}

	return newTable, h.TableRepository.AddSlice(oneTableList, helpers.NowTime())
}

type Handler struct {
	TableRepository table.Repository
}

func New(tables table.Repository) *Handler {
	return &Handler{tables}
}
