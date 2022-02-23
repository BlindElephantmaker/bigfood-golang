package createMass

import (
	"bigfood/internal/helpers"
)

type Message struct {
	CafeId   helpers.Uuid `json:"cafe-id" binding:"required" example:"uuid"`
	Quantity int          `json:"quantity" binding:"required" example:"20"`
}
