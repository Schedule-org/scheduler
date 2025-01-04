package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Services struct {
	gorm.Model
	ID             uuid.UUID    `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Name           string       `gorm:"size:100;not null"`
	Value          string       `gorm:"size:100;not null"`
	Duration       string       `gorm:"size:100;not null"`
	ProfessionalId uuid.UUID    `gorm:"type:uuid;not null"`
	Professional   Professional `gorm:"foreignKey:ProfessionalId"`
	CreatedAt      time.Time    `json:"created_at"`
	UpdatedAt      time.Time    `json:"updated_at"`
}
