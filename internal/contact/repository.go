package contact

import (
	"bigfood/internal/helpers"
	"time"
)

const (
	tableContact      = "contact"
	tableContactPhone = "contact_phone"
)

var NotExist = helpers.ErrorUnprocessableEntity("contact not exist")

type Repository interface {
	Add(contact Contact, createdAt time.Time) error
	GetByPhone(*helpers.Phone) (*Phone, error)
}
