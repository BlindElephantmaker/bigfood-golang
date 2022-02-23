package table

import (
	"bigfood/internal/helpers"
	"time"
)

type Repository interface {
	Add(tables []*Table, createdAt time.Time) error
	Get(tableId helpers.Uuid) (*Table, error)
	Update(*Table) error
	GetByCafe(cafeId helpers.Uuid) ([]*Table, error)
}
