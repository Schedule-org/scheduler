package models

import "time"

type Services struct {
	ID             string       `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Name           string       `gorm:"size:100;not null"`
	Value          string       `gorm:"size:100;not null"`
	Duration       string       `gorm:"size:100;not null"`
	ProfessionalId string       `gorm:"type:uuid;not null"`
	Professional   Professional `gorm:"foreignKey:ProfessionalId"`
	CreatedAt      time.Time    `json:"created_at"`
	UpdatedAt      time.Time    `json:"updated_at"`
}
