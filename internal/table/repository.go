package table

import (
	"bigfood/internal/cafe"
	"time"
)

const table = "tables"

type Repository interface {
	AddSlice(tables []*Table, createdAt time.Time) error
	Get(Id) (*Table, error)
	GetByCafe(cafe.Id) ([]*Table, error)
	Update(*Table) error
	Delete(Id) error
	DeleteAll(cafe.Id) error
}
