package cafe

import (
	"bigfood/internal/cafe/cafeUser"
	"time"
)

type Repository interface {
	Add(cafe *Cafe, cafeUser *cafeUser.User, createAt *time.Time) error
}
