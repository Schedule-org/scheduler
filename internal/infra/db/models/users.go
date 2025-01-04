package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	ID          uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Name        string    `gorm:"size:100;not null"`
	Email       string    `gorm:"size:100;unique;not null"`
	Password    string    `gorm:"not null"`
	ActivatedAt sql.NullTime
	CreatedAt   time.Time      `gorm:"autoCreateTime"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`

	// Relacionamento: Um Usuário pode ter muitos Estabelecimentos
	Establishments []Establishment `gorm:"foreignKey:UserID"`
}
