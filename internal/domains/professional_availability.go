package domains

import (
	"time"

	"github.com/google/uuid"
)

type ProfessionalAvailability struct {
	ProfessionalID uuid.UUID `json:"professional_id"`
	DayOfWeek      string    `json:"day_of_week"`
	StartTime      time.Time `json:"start_time"`
	EndTime        time.Time `json:"end_time"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
