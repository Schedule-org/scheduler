package db

import (
	"fmt"
	"time"

	"github.com/hebertzin/scheduler/internal/infra/db/models"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Migrate(database *gorm.DB) error {
	_ = logger.New(
		logrus.New(),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)
	err := database.AutoMigrate(&models.User{})

	if err != nil {
		return fmt.Errorf("failed to migrate models: %w", err)
	}

	fmt.Println("Database migrated successfully")
	return nil
}
