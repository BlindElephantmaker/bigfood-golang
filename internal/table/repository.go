package table

import (
	"bigfood/internal/helpers"
	"time"
)

type Repository interface {
	Add(tables []*Table, createdAt time.Time) error
	Get(cafeId helpers.Uuid) ([]*Table, error)
}
