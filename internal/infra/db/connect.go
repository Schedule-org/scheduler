package db

import (
	"fmt"
	"log"
	"time"

	"github.com/hebertzin/scheduler/internal/domain"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func ConnectDatabase(cfg *domain.Config) *gorm.DB {
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s port=%s host=%s",
		cfg.User, cfg.Password, cfg.Database, cfg.Port, cfg.Host)

	var err error

	gormLogger := logger.New(
		logrus.New(),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: gormLogger,
	})
	if err != nil {
		log.Fatalf("Failed to connect to the db: %v", err)
	}

	fmt.Println("Database connection successfully established")

	return db
}
