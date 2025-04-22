package factory

import (
	"github.com/hebertzin/scheduler/internal/core/usecases"
	"github.com/hebertzin/scheduler/internal/domain"

	"github.com/hebertzin/scheduler/internal/infra/db/repository"
	"github.com/hebertzin/scheduler/internal/presentation/controllers/v1"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func UsersFactory(db *gorm.DB, logger *logrus.Logger) domain.UserController {
	repo := repository.NewUserRepository(db, logger)
	useCase := usecases.NewAddUserUseCase(repo, logger)
	handler := controllers.NewUserController(useCase)
	return handler
}
