package reserve

import (
	"bigfood/internal/helpers"
	"time"
)

const tableReserve = "reserve"

var notExist = helpers.ErrorUnprocessableEntity("reserve not exist")
var errorReservedTimeIsBusy = helpers.ErrorBadRequest("table already reserved")

type Repository interface {
	Add(reserve *Reserve, createdAt time.Time) error
	Get(Id) (*Reserve, error)
	Update(*Reserve) error
	Delete(Id) error
	Undelete(Id) error
}
