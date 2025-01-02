package models

import "time"

type Professional struct {
	ID              string        `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Name            string        `gorm:"size:100;not null"`
	Role            string        `gorm:"size:100;not null"`
	EstablishmentId string        `gorm:"size:100;not null"`
	Establishment   Establishment `gorm:"foreignKey:EstablishmentId"`
	Services        []Services    `gorm:"foreignKey:ProfessionalId"`
	CreatedAt       time.Time     `json:"created_at"`
	UpdatedAt       time.Time     `json:"updated_at"`
}
