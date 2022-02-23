package table

import "time"

type Repository interface {
	Add(tables []*Table, createdAt time.Time) error
}
