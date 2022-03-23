package table

import (
	"bigfood/internal/helpers"
	"time"
)

type Repository interface {
	AddSlice(tables []*Table, createdAt time.Time) error
	Get(Id) (*Table, error)
	Update(*Table) error
	Delete(Id) error
	DeleteAll(cafeId helpers.Uuid) error
	GetByCafe(cafeId helpers.Uuid) ([]*Table, error)
}
