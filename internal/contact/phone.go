package contact

import (
	"bigfood/internal/helpers"
)

type Phone struct {
	ContactId Id            `json:"contact-id" db:"contact_id"`
	Phone     helpers.Phone `json:"phone" db:"phone"`
}

func NewPhone(phoneNumber helpers.Phone) *Phone {
	return &Phone{
		ContactId: newId(),
		Phone:     phoneNumber,
	}
}

func (p *Phone) GetId() Id {
	return p.ContactId
}

func (p *Phone) GetType() Type {
	return TypePhone
}
