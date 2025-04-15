package domains

import (
	"context"
	"time"

	"github.com/hebertzin/scheduler/internal/core"
)

type Services struct {
	Name           string    `json:"name"`
	Value          string    `json:"value"`
	Duration       string    `json:"duration"`
	ProfessionalId string    `json:"professional_id"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type ServicesUseCase interface {
	Add(ctx context.Context, payload *Services) (*Services, *core.Exception)

	FindServiceById(ctx context.Context, id string) (*Services, *core.Exception)

	GetAllServicesByProfessionalId(ctx context.Context, professional_id string) ([]Services, *core.Exception)
}

type ServicesRepository interface {
	Add(ctx context.Context, establishment *Services) (*Services, error)

	FindServiceById(ctx context.Context, service_id string) (*Services, error)

	GetAllServicesByProfessionalId(ctx context.Context, professional_id string) ([]Services, error)
}
