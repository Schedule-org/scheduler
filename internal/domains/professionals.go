package domains

import (
	"context"
	"time"

	"github.com/hebertzin/scheduler/internal/core"
)

type Professionals struct {
	Name            string    `json:"name"`
	Role            string    `json:"role"`
	EstablishmentId string    `json:"establishment"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type ProfessionalsUseCase interface {
	Add(ctx context.Context, payload *Professionals) (*Professionals, *core.Exception)
	FindProfessionalById(ctx context.Context, id string) (*Professionals, *core.Exception)
	UpdateProfessionalById(ctx context.Context, professional_id string, professionalData *Professionals) (*Professionals, *core.Exception)
}
