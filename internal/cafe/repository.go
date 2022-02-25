package cafe

import (
	"bigfood/internal/cafeUser"
)

type Repository interface {
	Add(cafe *Cafe, cafeUser *cafeUser.User) error
}
