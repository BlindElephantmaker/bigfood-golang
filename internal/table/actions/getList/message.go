package getList

import (
	"bigfood/internal/helpers"
)

type Message struct {
	CafeId helpers.Uuid `json:"cafe-id" binding:"required" example:"uuid"`
}
