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
	Id         Id         `json:"id" db:"id" example:"uuid"`
	Comment    Comment    `json:"comment" db:"comment"`
	GuestCount GuestCount `json:"guest-count" db:"guest_count" example:"4"`
	TableId    table.Id   `json:"table-id" db:"table_id" example:"uuid"`
	ContactId  contact.Id `json:"contact-id" db:"contact_id" example:"uuid"`
	FromDate   time.Time  `json:"from-date" db:"from_date" example:"RFC3339"`
	UntilDate  time.Time  `json:"until-date" db:"until_date" example:"RFC3339"`
	DeletedAt  *time.Time `json:"deleted-at,omitempty" db:"deleted_at" example:"RFC3339"`
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
