package domains

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hebertzin/scheduler/internal/core"
)

type (
	Establishment struct {
		Name       string    `json:"name"`
		City       string    `json:"city"`
		State      string    `json:"state"`
		PostalCode string    `json:"postal_code"`
		Number     string    `json:"number"`
		UserId     string    `json:"user_id"`
		CreatedAt  time.Time `json:"created_at"`
		UpdatedAt  time.Time `json:"updated_at"`
	}

	EstablishmentUseCase interface {
		Add(ctx context.Context, payload *Establishment) (*Establishment, *core.Exception)

		GetAllProfessionalsByEstablishmentId(ctx context.Context, establishment_id string) ([]Professionals, *core.Exception)

		FindEstablishmentById(ctx context.Context, establishment_id string) (*Establishment, *core.Exception)

		GetEstablishmentReport(ctx context.Context, establishment_id string) (*EstablishmentReport, *core.Exception)

		UpdateEstablishmentById(ctx context.Context, establishment_id string, establishmentData *Establishment) (*Establishment, *core.Exception)
	}

	EstablishmentRepository interface {
		Add(ctx context.Context, establishment *Establishment) (*Establishment, error)

		GetAllProfessionalsByEstablishmentId(ctx context.Context, establishment_id string) ([]Professionals, error)

		FindEstablishmentById(ctx context.Context, email string) (*Establishment, error)

		GetEstablishmentReport(ctx context.Context, establishment_id string) (*EstablishmentReport, error)

		UpdateEstablishmentById(ctx context.Context, establishment_id string, establishmentData *Establishment) (*Establishment, error)
	}

	EstablishmentController interface {
		Add(ctx *gin.Context)

		FindEstablishmentById(ctx *gin.Context)

		GetAllProfessinalsByEstablishmentId(ctx *gin.Context)

		UpdateEstablishmentById(ctx *gin.Context)

		GetEstablishmentReport(ctx *gin.Context)
	}
)
