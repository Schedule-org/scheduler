package domains

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/hebertzin/scheduler/internal/core"
)

type (
	ProfessionalAvailability struct {
		ProfessionalID uuid.UUID `json:"professional_id"`
		DayOfWeek      string    `json:"day_of_week"`
		StartTime      time.Time `json:"start_time"`
		EndTime        time.Time `json:"end_time"`
		CreatedAt      time.Time `json:"created_at"`
		UpdatedAt      time.Time `json:"updated_at"`
	}

	ProfessionalsAvailabilityUseCase interface {
		Add(ctx context.Context, availability *ProfessionalAvailability) (*ProfessionalAvailability, *core.Exception)

		GetProfessionalAvailabilityById(ctx context.Context, professional_id string) ([]ProfessionalAvailability, *core.Exception)
	}

	ProfessionalsAvailabilityRepository interface {
		Add(ctx context.Context, availability *ProfessionalAvailability) (*ProfessionalAvailability, error)

		GetProfessionalAvailabilityById(ctx context.Context, professional_id string) ([]ProfessionalAvailability, error)
	}

	ProfessionalAvailabilityController interface {
		Add(ctx *gin.Context)

		GetProfessionalAvailabilityById(ctx *gin.Context)
	}
)
