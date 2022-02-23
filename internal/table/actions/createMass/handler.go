package createMass

import (
	"bigfood/internal/helpers"
	"bigfood/internal/table"
	"errors"
	"fmt"
)

var ErrorQuantityIsTooLow = errors.New("quantity of tables must be greater than 0")
var ErrorQuantityIsTooHigh = errors.New("quantity of tables must be less than 100")

func (h *Handler) Run(message *Message) ([]*table.Table, error) {
	if message.Quantity < 1 {
		return nil, ErrorQuantityIsTooLow
	}
	if message.Quantity > 100 {
		return nil, ErrorQuantityIsTooHigh
	}

	var tables []*table.Table
	for i := 1; i <= message.Quantity; i++ {
		title, _ := table.NewTitle(fmt.Sprint(i))
		tables = append(tables, table.NewTable(message.CafeId, title))
	}

	err := h.TableRepository.Add(tables, helpers.TimeNow())
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
