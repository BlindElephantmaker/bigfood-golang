package table

import "errors"

type Seats int

var ErrorTableSeatsInvalidValue = errors.New("table seats has invalid value")

func NewSeats() Seats {
	return 4
}

func ParseSeats(seats int) (Seats, error) {
	if seats < 1 {
		return 0, ErrorTableSeatsInvalidValue
	}

	return Seats(seats), nil
}
