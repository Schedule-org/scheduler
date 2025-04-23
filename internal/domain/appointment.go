package domain

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hebertzin/scheduler/internal/core"
)

type (
	Appointment struct {
		ProfessionalId string    `json:"professional_id"`
		ServiceId      string    `json:"service_id"`
		ScheduledDate  time.Time `json:"schedule_date"`
		Email          string    `json:"user_email"`
		Phone          string    `json:"user_phone"`
		Notes          string    `json:"notes"`
		CreatedAt      time.Time `json:"created_at"`
		UpdatedAt      time.Time `json:"updated_at"`
	}

	// AppointmentUseCase defines the business logic for appointments.
	AppointmentUseCase interface {
		// Add creates a new appointment.
		// Returns the created appointment and a possible business exception.
		Add(ctx context.Context, appointment *Appointment) (*Appointment, *core.Exception)
		// GetAllAppointmentsByProfessionalId retrieves all appointments for a given professional.
		// Returns a list of appointments and a possible business exception.
		GetAllAppointmentsByProfessionalId(ctx context.Context, professional_id string) ([]Appointment, *core.Exception)
		// GetAppointmentById retrieves a specific appointment by its ID.
		// Returns the appointment and a possible business exception.
		GetAppointmentById(ctx context.Context, appointment_id string) (*Appointment, *core.Exception)
	}

	// AppointmentRepository defines the data persistence methods for appointments.
	AppointmentRepository interface {
		// Add inserts a new appointment into the database.
		// Returns the saved appointment and a possible error.
		Add(ctx context.Context, appointment *Appointment) (*Appointment, error)
		// GetAllAppointmentsByProfessionalId fetches all appointments for a specific professional from the database.
		// Returns a list of appointments and a possible error.
		GetAllAppointmentsByProfessionalId(ctx context.Context, professional_id string) ([]Appointment, error)
		// GetAppointmentById fetches an appointment by its ID from the database.
		// Returns the appointment and a possible error.
		GetAppointmentById(ctx context.Context, appointment_id string) (*Appointment, error)

		DeleteAppointment(ctx context.Context, appointment_id string) error
	}

	// AppointmentController defines the HTTP handlers for appointments.
	AppointmentController interface {
		// Add handles the HTTP request to create a new appointment.
		Add(ctx *gin.Context)
		// GetAllAppointmentsByProfessionalId handles the HTTP request to fetch all appointments for a professional.
		GetAllAppointmentsByProfessionalId(ctx *gin.Context)
		// GetAppointmentById handles the HTTP request to fetch a specific appointment by its ID.
		GetAppointmentById(ctx *gin.Context)
	}
)
