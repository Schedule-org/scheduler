package domains

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hebertzin/scheduler/internal/core"
)

type (
	Appointment struct {
		ProfessionalID string    `json:"professional_id"`
		ServiceID      string    `json:"service_id"`
		ScheduledDate  time.Time `json:"schedule_date"`
		Email          string    `json:"user_email"`
		Phone          string    `json:"user_phone"`
		Notes          string    `json:"notes"`
		CreatedAt      time.Time `json:"created_at"`
		UpdatedAt      time.Time `json:"updated_at"`
	}

	AppointmentUseCase interface {
		Add(ctx context.Context, appointment *Appointment) (*Appointment, *core.Exception)

		GetAllAppointmentsByProfessionalId(ctx context.Context, professional_id string) ([]Appointment, *core.Exception)

		GetAppointmentById(ctx context.Context, appointment_id string) (*Appointment, *core.Exception)
	}

	AppointmentRepository interface {
		Add(ctx context.Context, appointment *Appointment) (*Appointment, error)

		GetAllAppointmentsByProfessionalId(ctx context.Context, professional_id string) ([]Appointment, error)

		GetAppointmentById(ctx context.Context, appointment_id string) (*Appointment, error)
	}

	AppointmentController interface {
		Add(ctx *gin.Context)

		GetAllAppointmentsByProfessionalId(ctx *gin.Context)

		GetAppointmentById(ctx *gin.Context)
	}
)
