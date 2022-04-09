package reserve

import (
	"bigfood/internal/helpers"
	"bigfood/internal/table"
	"time"
)

const tableReserve = "reserve"

var notExist = helpers.ErrorUnprocessableEntity("reserve not exist")
var errorReservedTimeIsBusy = helpers.ErrorBadRequest("table already reserved")

type Repository interface {
	Add(reserve *Reserve, createdAt time.Time) error
	Get(Id) (*Reserve, error)
	GetActualByTableId(table.Id) ([]*Reserve, error)
	GetDeletedByTableId(table.Id) ([]*Reserve, error)
	GetHistoryByTableId(tableId table.Id, limit int, offset int) ([]*Reserve, error)
	Update(*Reserve) error
	Delete(Id) error
	Undelete(Id) error
}
