package createMass

import (
	"bigfood/internal/helpers"
	"bigfood/internal/table"
	"fmt"
)

type Message struct {
	CafeId   helpers.Uuid `json:"cafe-id" binding:"required" example:"uuid"`
	Quantity Quantity     `json:"quantity" binding:"required" example:"10"`
}

func (h *Handler) Run(message *Message) ([]*table.Table, error) {
	var tables []*table.Table
	for i := 1; i <= int(message.Quantity); i++ {
		title, _ := table.ParseTitle(fmt.Sprint(i))
		tables = append(tables, table.NewTable(message.CafeId, title))
	}

	err := h.TableRepository.AddSlice(tables, helpers.TimeNow())
	if err != nil {
		return nil, err
	}

	return tables, nil
}

type Handler struct {
	TableRepository table.Repository
}

func New(tables table.Repository) *Handler {
	return &Handler{tables}
}
