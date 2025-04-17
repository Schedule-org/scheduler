package domains

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hebertzin/scheduler/internal/core"
)

type (
	Professionals struct {
		Name            string    `json:"name"`
		Role            string    `json:"role"`
		EstablishmentId string    `json:"establishment"`
		CreatedAt       time.Time `json:"created_at"`
		UpdatedAt       time.Time `json:"updated_at"`
	}

	// ProfessionalsUseCase defines the business logic for managing professionals.
	ProfessionalsUseCase interface {
		// Add creates a new professional.
		Add(ctx context.Context, payload *Professionals) (*Professionals, *core.Exception)
		// FindProfessionalById retrieves a professional by their ID.
		FindProfessionalById(ctx context.Context, id string) (*Professionals, *core.Exception)
		// UpdateProfessionalById updates the data of a professional by their ID.
		UpdateProfessionalById(ctx context.Context, professional_id string, professionalData *Professionals) (*Professionals, *core.Exception)
	}

	// ProfessionalsRepository defines the data access layer for professionals.
	ProfessionalsRepository interface {
		// Add inserts a new professional into the database.
		Add(ctx context.Context, establishment *Professionals) (*Professionals, error)
		// FindProfessionalById retrieves a professional by their ID from the database.
		FindProfessionalById(ctx context.Context, email string) (*Professionals, error)
		// UpdateProfessionalById updates the data of an existing professional in the database.
		UpdateProfessionalById(ctx context.Context, professional_id string, professionalData *Professionals) (*Professionals, error)
	}

	// ProfessionalsController defines the HTTP handlers for professional-related operations.
	ProfessionalsController interface {
		// Add handles the HTTP request to create a new professional.
		Add(ctx *gin.Context)
		// FindProfessionalById handles the HTTP request to retrieve a professional by ID.
		FindProfessionalById(ctx *gin.Context)
		// UpdateProfessionalById handles the HTTP request to update a professional's information by ID.
		UpdateProfessionalById(ctx *gin.Context)
	}
)
