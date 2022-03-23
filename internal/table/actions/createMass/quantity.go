package createMass

import (
	"bigfood/internal/helpers"
	"encoding/json"
)

type Quantity int

var errorQuantityIsTooLow = helpers.NewErrorBadRequest("quantity of tables must be greater than 0")
var errorQuantityIsTooHigh = helpers.NewErrorBadRequest("quantity of tables must be less than 100")

func (q *Quantity) UnmarshalJSON(data []byte) error {
	var value int
	err := json.Unmarshal(data, &value)
	if err != nil {
		return err
	}

	quantity, err := parseQuantity(value)
	if err != nil {
		return err
	}

	*q = quantity
	return nil
}

func parseQuantity(value int) (Quantity, error) {
	if value < 1 {
		return 0, errorQuantityIsTooLow
	}
	if value > 100 {
		return 0, errorQuantityIsTooHigh
	}

	return Quantity(value), nil
}
