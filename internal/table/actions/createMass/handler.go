package createMass

import (
	"bigfood/internal/cafe"
	"bigfood/internal/helpers"
	"bigfood/internal/table"
	"fmt"
)

type Message struct {
	CafeId   cafe.Id  `json:"cafe-id" binding:"required" example:"uuid"`
	Quantity Quantity `json:"quantity" binding:"required" example:"10"`
}

func (h *Handler) Run(m *Message) ([]*table.Table, error) {
	var tables []*table.Table
	for i := 1; i <= int(m.Quantity); i++ {
		title, _ := table.ParseTitle(fmt.Sprint(i))
		tables = append(tables, table.NewTable(m.CafeId, title))
	}

	err := h.TableRepository.AddSlice(tables, helpers.NowTime())
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
