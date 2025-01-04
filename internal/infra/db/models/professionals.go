package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Professional struct {
	gorm.Model
	ID              uuid.UUID     `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Name            string        `gorm:"size:100;not null"`
	Role            string        `gorm:"size:100;not null"`
	EstablishmentId uuid.UUID     `gorm:"type:uuid;not null"`
	Establishment   Establishment `gorm:"foreignKey:EstablishmentId"`
	Services        []Services    `gorm:"foreignKey:ProfessionalId"`
	CreatedAt       time.Time     `json:"created_at"`
	UpdatedAt       time.Time     `json:"updated_at"`
}
