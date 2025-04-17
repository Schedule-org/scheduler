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
		ProfessionalId uuid.UUID `json:"professional_id"`
		DayOfWeek      string    `json:"day_of_week"`
		StartTime      time.Time `json:"start_time"`
		EndTime        time.Time `json:"end_time"`
		CreatedAt      time.Time `json:"created_at"`
		UpdatedAt      time.Time `json:"updated_at"`
	}

	// ProfessionalsAvailabilityUseCase defines the business logic related to professional availability.
	ProfessionalsAvailabilityUseCase interface {
		// Add creates a new availability record for a professional.
		Add(ctx context.Context, availability *ProfessionalAvailability) (*ProfessionalAvailability, *core.Exception)
		// GetProfessionalAvailabilityById retrieves all availability records for a given professional.
		GetProfessionalAvailabilityById(ctx context.Context, professional_id string) ([]ProfessionalAvailability, *core.Exception)
	}

	// ProfessionalsAvailabilityRepository defines the data access layer for professional availability.
	ProfessionalsAvailabilityRepository interface {
		// Add inserts a new availability record into the database.
		Add(ctx context.Context, availability *ProfessionalAvailability) (*ProfessionalAvailability, error)
		// GetProfessionalAvailabilityById fetches all availability records for a professional from the database.
		GetProfessionalAvailabilityById(ctx context.Context, professional_id string) ([]ProfessionalAvailability, error)
	}

	// ProfessionalAvailabilityController defines the HTTP handlers for managing professional availability.
	ProfessionalAvailabilityController interface {
		// Add handles the HTTP request to create a new availability entry for a professional.
		Add(ctx *gin.Context)
		// GetProfessionalAvailabilityById handles the HTTP request to retrieve a professional's availability by ID.
		GetProfessionalAvailabilityById(ctx *gin.Context)
	}
)
