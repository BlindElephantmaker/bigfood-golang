package tableDelete

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

	return h.TableRepository.Delete(tableId)
}
