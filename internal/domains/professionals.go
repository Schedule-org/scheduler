package domains

import "time"

type Professionals struct {
	Name            string    `json:"name"`
	Role            string    `json:"role"`
	EstablishmentId string    `json:"establishment"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
