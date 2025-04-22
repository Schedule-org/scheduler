package domain

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hebertzin/scheduler/internal/core"
)

type (
	Services struct {
		Name           string    `json:"name"`
		Value          string    `json:"value"`
		Duration       string    `json:"duration"`
		ProfessionalId string    `json:"professional_id"`
		CreatedAt      time.Time `json:"created_at"`
		UpdatedAt      time.Time `json:"updated_at"`
	}

	// ServicesUseCase defines the business logic for managing services.
	ServicesUseCase interface {
		// Add creates a new service.
		Add(ctx context.Context, payload *Services) (*Services, *core.Exception)
		// FindServiceById retrieves a service by its ID.
		FindServiceById(ctx context.Context, id string) (*Services, *core.Exception)
		// GetAllServicesByProfessionalId returns all services associated with a specific professional.
		GetAllServicesByProfessionalId(ctx context.Context, professional_id string) ([]Services, *core.Exception)
	}

	// ServicesRepository defines the data access layer for services.
	ServicesRepository interface {
		// Add inserts a new service into the database.
		Add(ctx context.Context, establishment *Services) (*Services, error)
		// FindServiceById retrieves a service by its ID from the database.
		FindServiceById(ctx context.Context, service_id string) (*Services, error)
		// GetAllServicesByProfessionalId fetches all services linked to a specific professional from the database.
		GetAllServicesByProfessionalId(ctx context.Context, professional_id string) ([]Services, error)
	}

	// ServicesController defines the HTTP handlers for managing services.
	ServicesController interface {
		// Add handles the HTTP request to create a new service.
		Add(ctx *gin.Context)
		// FindServiceById handles the HTTP request to retrieve a service by ID.
		FindServiceById(ctx *gin.Context)
		// GetAllServicesByProfessionalId handles the HTTP request to list all services for a given professional.
		GetAllServicesByProfessionalId(ctx *gin.Context)
	}
)
