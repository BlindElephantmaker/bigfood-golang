package createOrganization

import (
	"bigfood/internal/organization"
	"github.com/google/uuid"
)

type Handler struct {
	organizationRepository organization.Repository
}

func New(organizations organization.Repository) *Handler {
	return &Handler{
		organizationRepository: organizations,
	}
}

func (h *Handler) Run(userId *uuid.UUID) (*uuid.UUID, error) {
	org := organization.New()
	orgUser := organization.NewOrganizationUser(org.Id, userId)
	err := h.organizationRepository.Add(org, orgUser)
	if err != nil {
		return nil, err
	}

	return org.Id, nil
}
