package db

import (
	"fmt"

	"github.com/hebertzin/scheduler/internal/infra/db/models"
	"gorm.io/gorm"
)

func Migrate(database *gorm.DB) error {
	err := database.AutoMigrate(&models.Users{})
	if err != nil {
		return fmt.Errorf("failed to migrate models: %w", err)
	}
	fmt.Println("Database migrated successfully")
	return nil
}
