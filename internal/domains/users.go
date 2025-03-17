package domains

import (
	"context"
	"database/sql"
	"time"

	"github.com/hebertzin/scheduler/internal/core"
)

type User struct {
	Name        string       `json:"name"`
	Email       string       `json:"email"`
	Password    string       `json:"password"`
	ActivatedAt sql.NullTime `json:"activate_at"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
}

type UserUseCase interface {
	Add(ctx context.Context, payload *User) (*User, *core.Exception)
	FindUserById(ctx context.Context, id string) (*User, *core.Exception)
	FindAllUsers(ctx context.Context) ([]User, *core.Exception)
	FindAllEstablishmentsByUserId(ctx context.Context, user_id string) ([]Establishment, *core.Exception)
}
