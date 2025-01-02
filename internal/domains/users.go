package domains

import (
	"database/sql"
	"time"
)

type User struct {
	Name        string       `json:"name"`
	Email       string       `json:"email"`
	Password    string       `json:"password"`
	ActivatedAt sql.NullTime `json:"activate_at"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
}
