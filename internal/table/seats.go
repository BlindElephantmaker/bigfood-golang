package table

import (
	"bigfood/internal/helpers"
	"encoding/json"
)

type Seats int

const defaultSeatsValue = Seats(4)

var errorTableSeatsInvalidValue = helpers.NewErrorBadRequest("table seats has invalid value")

func (s *Seats) UnmarshalJSON(data []byte) error {
	var value int
	err := json.Unmarshal(data, &value)
	if err != nil {
		return err
	}

	seats, err := parseSeats(value)
	if err != nil {
		return err
	}

	*s = seats
	return nil
}

func NewSeats() Seats {
	return defaultSeatsValue
}

func parseSeats(seats int) (Seats, error) {
	if seats < 1 {
		return 0, errorTableSeatsInvalidValue
	}

	return Seats(seats), nil
}
