package organization

import (
	"bigfood/internal/helpers/timeHelper"
	"github.com/google/uuid"
	"time"
)

type Organization struct {
	Id        *uuid.UUID
	CreatedAt *time.Time
}

func New() *Organization {
	id := uuid.New()
	now := timeHelper.Now()
	return &Organization{
		Id:        &id,
		CreatedAt: &now,
	}
}
