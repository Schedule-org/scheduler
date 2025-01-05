package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProfessionalAvailability struct {
	gorm.Model
	ID             uuid.UUID    `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	ProfessionalID uuid.UUID    `gorm:"type:uuid;not null"`
	Professional   Professional `gorm:"foreignKey:ProfessionalID"`
	DayOfWeek      string       `gorm:"size:20;not null"`
	StartTime      time.Time    `gorm:"not null"`
	EndTime        time.Time    `gorm:"not null"`
	CreatedAt      time.Time    `json:"created_at"`
	UpdatedAt      time.Time    `json:"updated_at"`
}
