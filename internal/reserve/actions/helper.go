package reserveAction

import (
	"bigfood/internal/contact"
	"bigfood/internal/helpers"
	"bigfood/internal/reserve"
	"bigfood/internal/table"
	"time"
)

type Helper struct {
	contactRepository contact.Repository
	tableRepository   table.Repository
}

func NewHelper(
	contactRepository contact.Repository,
	tableRepository table.Repository,
) *Helper {
	return &Helper{
		contactRepository: contactRepository,
		tableRepository:   tableRepository,
	}
}

type Message struct {
	TableId    table.Id           `json:"table-id" binding:"required" example:"uuid"`
	Phone      *helpers.Phone     `json:"phone" example:"+71234567890"`
	ContactId  *contact.Id        `json:"contact-id" example:"uuid"`
	GuestCount reserve.GuestCount `json:"guest-count" binding:"required" example:"4"`
	Comment    reserve.Comment    `json:"comment" binding:"required"`
	FromData   time.Time          `json:"from-data" binding:"required" example:"RFC3339"`
	UntilData  time.Time          `json:"until-data" binding:"required" example:"RFC3339"`
}

var errorPhoneOrContactIdMustBePassed = helpers.ErrorBadRequest("phone or contact_id must be passed")
var errorGuestCountMustBeLessThanTableSeatsCount = helpers.ErrorBadRequest("guest count must be less than table seats count")

func (h *Helper) ParseMessage(id reserve.Id, m *Message) (*reserve.Reserve, error) {
	if m.ContactId == nil && m.Phone == nil {
		return nil, errorPhoneOrContactIdMustBePassed
	}

	t, err := h.tableRepository.Get(m.TableId)
	if err != nil {
		return nil, err
	}
	if int(m.GuestCount) > int(t.Seats) {
		return nil, errorGuestCountMustBeLessThanTableSeatsCount
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

	return reserve.NewReserve(
		id,
		m.TableId,
		contactId,
		m.Comment,
		m.GuestCount,
		m.FromData,
		m.UntilData,
	)
}
