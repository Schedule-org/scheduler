package domains

import "time"

type Services struct {
	Name           string    `json:"name"`
	Value          string    `json:"value"`
	Duration       string    `json:"duration"`
	ProfessionalId string    `json:"professional_id"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
