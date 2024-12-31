package env

import (
	"os"

	"github.com/hebertzin/scheduler/internal/domains"
	"github.com/hebertzin/scheduler/internal/infra/config/logging"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func LoadConfig() *domains.Config {
	err := godotenv.Load()
	if err != nil {
		logging.Log.WithFields(logrus.Fields{
			"method": "LoadConfig",
		}).Error("Some error has been ocurred load env variables")
	}

	config := &domains.Config{
		User:     os.Getenv("USER_DATABASE"),
		Password: os.Getenv("USER_PASSWORD"),
		Database: os.Getenv("DATABASE"),
		Port:     os.Getenv("PORT"),
		Host:     os.Getenv("HOST"),
	}

	if config.User == "" || config.Password == "" || config.Database == "" || config.Port == "" || config.Host == "" {
		logging.Log.WithFields(logrus.Fields{
			"method":   "LoadConfig",
			"user":     config.User,
			"database": config.Database,
		}).Error("Some error has been ocurred load env variables")
	}

	return config
}
