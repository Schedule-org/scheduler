package dto

import (
	"database/sql"
	"time"

	"github.com/hebertzin/scheduler/internal/domains"
)

type UserDTO struct {
	Name        string       `json:"name"`
	Email       string       `json:"email"`
	ActivatedAt sql.NullTime `json:"activate_at"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
}

func MapToUserDTO(user *domains.User) *UserDTO {
	return &UserDTO{
		Name:        user.Name,
		Email:       user.Email,
		ActivatedAt: user.ActivatedAt,
		CreatedAt:   user.CreatedAt,
		UpdatedAt:   user.UpdatedAt,
	}
}
