package tableEdit

import (
	"bigfood/internal/helpers"
	"bigfood/internal/table"
)

type Handler struct {
	TableRepository table.Repository
}

func New(tables table.Repository) *Handler {
	return &Handler{tables}
}

func (h *Handler) Run(message *Message) error {
	tableId, err := helpers.UuidParse(message.TableId)
	if err != nil {
		return err
	}
	t, err := h.TableRepository.Get(tableId)
	if err != nil {
		return err
	}

	if message.Title != nil {
		title, err := table.ParseTitle(*message.Title)
		if err != nil {
			return err
		}
		t.Title = title
	}
	if message.Comment != nil {
		comment, err := table.ParseComment(*message.Comment)
		if err != nil {
			return err
		}
		t.Comment = comment
	}
	if message.Seats != nil {
		seats, err := table.ParseSeats(*message.Seats)
		if err != nil {
			return err
		}
		t.Seats = seats
	}

	return h.TableRepository.Update(t)
}
