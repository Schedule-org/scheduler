package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Establishment struct {
	gorm.Model
	ID         uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Name       string    `gorm:"size:100;not null"`
	City       string    `gorm:"size:100;not null"`
	State      string    `gorm:"size:100;not null"`
	PostalCode string    `gorm:"size:100;not null"`
	Number     string    `gorm:"size:100;not null"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`

	UserID   uuid.UUID `gorm:"type:uuid;not null"`
	Accounts Accounts  `gorm:"foreignKey:UserID"`
}
