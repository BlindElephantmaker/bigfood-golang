package tableCreate

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

func (h *Handler) Run(message *Message) (*table.Table, error) {
	newTable, err := parseMessageToTable(message)
	if err != nil {
		return nil, err
	}

	oneTableList := []*table.Table{newTable}

	return newTable, h.TableRepository.AddSlice(oneTableList, helpers.TimeNow())
}

func parseMessageToTable(message *Message) (*table.Table, error) {
	cafeId, err := helpers.UuidParse(message.CafeId)
	if err != nil {
		return nil, err
	}

	var title table.Title
	if message.Title != nil {
		title, err = table.ParseTitle(*message.Title)
		if err != nil {
			return nil, err
		}
	} else {
		title, _ = table.ParseTitle("New table") // todo: serial number
	}

	var comment table.Comment
	if message.Comment != nil {
		comment, err = table.ParseComment(*message.Comment)
		if err != nil {
			return nil, err
		}
	} else {
		comment, _ = table.ParseComment("")
	}

	var seats table.Seats
	if message.Seats != nil {
		seats, err = table.ParseSeats(*message.Seats)
		if err != nil {
			return nil, err
		}
	} else {
		seats = table.NewSeats()
	}

	return &table.Table{
		Id:      helpers.UuidGenerate(),
		CafeId:  cafeId,
		Title:   title,
		Comment: comment,
		Seats:   seats,
	}, nil
}
