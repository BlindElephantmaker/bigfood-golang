package cafe

import (
	"github.com/google/uuid"
)

type Cafe struct {
	Id *uuid.UUID
}

func New() *Cafe {
	id := uuid.New()
	return &Cafe{&id}
}
