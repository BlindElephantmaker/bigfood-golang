package cafe

import (
	"bigfood/internal/helpers"
)

type Cafe struct {
	Id helpers.Uuid
}

func New() *Cafe {
	return &Cafe{helpers.NewUuid()}
}
