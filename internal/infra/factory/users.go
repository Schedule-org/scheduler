package factory

import (
	"github.com/hebertzin/scheduler/internal/core/usecases"
	"github.com/hebertzin/scheduler/internal/infra/db/repository"
	"github.com/hebertzin/scheduler/internal/presentation/controllers/v1"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func UsersFactory(db *gorm.DB, logger *logrus.Logger) controllers.UserController {
	ur := repository.NewUserRepository(db, logger)
	uc := usecases.NewAddUserUseCase(ur, logger)
	uh := controllers.NewUserController(uc)
	return uh
}
