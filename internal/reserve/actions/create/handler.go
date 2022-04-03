package reserveCreate

import (
	"bigfood/internal/contact"
	"bigfood/internal/helpers"
	"bigfood/internal/reserve"
	"bigfood/internal/table"
	"time"
)

type Message struct {
	TableId    table.Id           `json:"table-id" binding:"required" example:"uuid"`
	Phone      *helpers.Phone     `json:"phone" example:"+71234567890"`
	ContactId  *contact.Id        `json:"contact-id" example:"uuid"`
	GuestCount reserve.GuestCount `json:"guest-count" binding:"required" example:"4"`
	Comment    reserve.Comment    `json:"comment" binding:"required"`
	FromData   time.Time          `json:"from-data" binding:"required" example:"RFC3339"`
	UntilData  time.Time          `json:"until-data" binding:"required" example:"RFC3339"`
}

var ErrorPhoneOrContactIdMustBePassed = helpers.ErrorBadRequest("phone or contact_id must be passed")
var ErrorGuestCountMustBeLessThanTableSeatsCount = helpers.ErrorBadRequest("guest count must be less than table seats count")

func (h *Handler) Run(m *Message) (*reserve.Reserve, error) {
	if m.ContactId == nil && m.Phone == nil {
		return nil, ErrorPhoneOrContactIdMustBePassed
	}

	t, err := h.tableRepository.Get(m.TableId)
	if err != nil {
		return nil, err
	}
	if int(m.GuestCount) > int(t.Seats) {
		return nil, ErrorGuestCountMustBeLessThanTableSeatsCount
	}

	var contactId contact.Id
	if m.ContactId != nil {
		contactId = *m.ContactId
	} else {
		var contactPhone *contact.Phone
		contactPhone, err = h.contactRepository.GetByPhone(m.Phone)
		if err == contact.NotExist {
			contactPhone = contact.NewPhone(*m.Phone)
			err = h.contactRepository.Add(contactPhone, helpers.NowTime())
		}
		if err != nil {
			return nil, err
		}
		contactId = contactPhone.GetId()
	}

	newReserve, err := reserve.NewReserve(
		m.TableId,
		contactId,
		m.Comment,
		m.GuestCount,
		m.FromData,
		m.UntilData,
	)
	if err != nil {
		return nil, err
	}

	if err := h.reserveRepository.Add(newReserve, helpers.NowTime()); err != nil {
		return nil, err
	}

	return newReserve, nil
}

type Handler struct {
	reserveRepository reserve.Repository
	contactRepository contact.Repository
	tableRepository   table.Repository
}

func New(
	reserves reserve.Repository,
	contacts contact.Repository,
	tables table.Repository,

) *Handler {
	return &Handler{
		reserveRepository: reserves,
		contactRepository: contacts,
		tableRepository:   tables,
	}
}
