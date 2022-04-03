package reserve

import (
	"bigfood/internal/contact"
	"bigfood/internal/helpers"
	"bigfood/internal/table"
	"time"
)

var ErrorFromDateMustBeLessThanUntilDate = helpers.ErrorBadRequest("from date must be less than until date")
var ErrorFromDateMustBeMoreThanCurrentDate = helpers.ErrorBadRequest("from date must be more than current date")

type Reserve struct {
	Id         Id         `json:"id" db:"id"`
	Comment    Comment    `json:"comment" db:"comment"`
	GuestCount GuestCount `json:"guest-count" db:"guest_count"`
	TableId    table.Id   `json:"table-id" db:"table_id"`
	ContactId  contact.Id `json:"contact-id" db:"contact_id"`
	FromDate   time.Time  `json:"from-date" db:"from_date"`
	UntilDate  time.Time  `json:"until-date" db:"until_date"`
	DeletedAt  *time.Time `json:"deleted-at,omitempty" db:"deleted_at"`
}

func NewReserve(
	tableId table.Id,
	contactId contact.Id,
	comment Comment,
	guestCount GuestCount,
	from, until time.Time,
) (*Reserve, error) {
	if !from.Before(until) {
		return nil, ErrorFromDateMustBeLessThanUntilDate
	}
	if from.Before(helpers.NowTime()) {
		return nil, ErrorFromDateMustBeMoreThanCurrentDate
	}

	return &Reserve{
		Id:         newId(),
		TableId:    tableId,
		ContactId:  contactId,
		Comment:    comment,
		GuestCount: guestCount,
		FromDate:   from,
		UntilDate:  until,
		DeletedAt:  nil,
	}, nil
}
