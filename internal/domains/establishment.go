package domains

import (
	"context"
	"time"

	"github.com/hebertzin/scheduler/internal/core"
)

type Establishment struct {
	Name       string    `json:"name"`
	City       string    `json:"city"`
	State      string    `json:"state"`
	PostalCode string    `json:"postal_code"`
	Number     string    `json:"number"`
	UserId     string    `json:"user_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type EstablishmentUseCase interface {
	Add(ctx context.Context, payload *Establishment) (*Establishment, *core.Exception)
	GetAllProfessionalsByEstablishmentId(ctx context.Context, establishment_id string) ([]Professionals, *core.Exception)
	FindEstablishmentById(ctx context.Context, establishment_id string) (*Establishment, *core.Exception)
	GetEstablishmentReport(ctx context.Context, establishment_id string) (*EstablishmentReport, *core.Exception)
	UpdateEstablishmentById(ctx context.Context, establishment_id string, establishmentData *Establishment) (*Establishment, *core.Exception)
}
