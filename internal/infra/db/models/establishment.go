package models

import (
	"time"
)

type Establishment struct {
	ID         string    `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Name       string    `gorm:"size:100;not null"`
	City       string    `gorm:"size:100;not null"`
	State      string    `gorm:"size:100;not null"`
	PostalCode string    `gorm:"size:100;not null"`
	Number     string    `gorm:"size:100;not null"`
	UserId     string    `gorm:"size:100;not null"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`

	User Users `gorm:"foreignKey:UserId"`
}
