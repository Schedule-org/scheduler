package domains

import (
	"time"

	"github.com/google/uuid"
)

type Appointment struct {
	ProfessionalID uuid.UUID `json:"professional_id"`
	ServiceID      uuid.UUID `json:"service_id"`
	ScheduledDate  time.Time `json:"schedule_date"`
	Email          string    `json:"user_email"`
	Phone          string    `json:"user_phone"`
	Notes          string    `json:"notes"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
