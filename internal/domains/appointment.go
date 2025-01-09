package domains

import (
	"time"
)

type Appointment struct {
	ProfessionalID string    `json:"professional_id"`
	ServiceID      string    `json:"service_id"`
	ScheduledDate  time.Time `json:"schedule_date"`
	Email          string    `json:"user_email"`
	Phone          string    `json:"user_phone"`
	Notes          string    `json:"notes"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
