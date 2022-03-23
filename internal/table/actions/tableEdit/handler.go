package tableEdit

import (
	"bigfood/internal/table"
)

type Message struct {
	TableId table.Id       `json:"table-id" binding:"required" example:"uuid"`
	Title   *table.Title   `json:"title"`
	Comment *table.Comment `json:"comment"`
	Seats   *table.Seats   `json:"seats"`
}

func (h *Handler) Run(m *Message) error {
	t, err := h.TableRepository.Get(m.TableId)
	if err != nil {
		return err
	}

	if m.Title != nil {
		t.Title = *m.Title
	}
	if m.Comment != nil {
		t.Comment = *m.Comment
	}
	if m.Seats != nil {
		t.Seats = *m.Seats
	}

	return h.TableRepository.Update(t)
}

type Handler struct {
	TableRepository table.Repository
}

func New(tables table.Repository) *Handler {
	return &Handler{tables}
}
