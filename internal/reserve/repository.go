package reserve

import (
	"bigfood/internal/helpers"
	"time"
)

const tableReserve = "reserve"

var errorReservedTimeIsBusy = helpers.ErrorBadRequest("table already reserved")

type Repository interface {
	Add(reserve *Reserve, createdAt time.Time) error
}
