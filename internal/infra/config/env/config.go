package env

import (
	"os"

	"github.com/hebertzin/scheduler/internal/domain"
	"github.com/hebertzin/scheduler/internal/infra/config/logging"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func LoadEnvConfig() *domain.Config {
	err := godotenv.Load()
	if err != nil {
		logging.Log.WithFields(logrus.Fields{
			"method": "LoadConfig",
		}).Error("Some error has been ocurred load env variables")
	}

	config := &domain.Config{
		User:     os.Getenv("USER_DATABASE"),
		Password: os.Getenv("USER_PASSWORD"),
		Database: os.Getenv("DATABASE"),
		Port:     os.Getenv("PORT"),
		Host:     os.Getenv("HOST"),
	}

	if config.User == "" || config.Password == "" || config.Database == "" || config.Port == "" || config.Host == "" {
		println("Some env fields are missing")
	}

	return config
}
