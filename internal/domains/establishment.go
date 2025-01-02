package domains

import "time"

type Establishment struct {
	Name       string    `json:"name"`
	City       string    `json:"city"`
	State      string    `json:"state"`
	PostalCode string    `json:"postal_code"`
	Number     string    `json:"number"`
	UserId     string    `json:"user_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
