package factory

import (
	"github.com/hebertzin/scheduler/internal/core/usecases"
	"github.com/hebertzin/scheduler/internal/domain"

	"github.com/hebertzin/scheduler/internal/infra/db/repository"
	"github.com/hebertzin/scheduler/internal/presentation/controllers"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func AccountFactory(db *gorm.DB, logger *logrus.Logger) domain.AccountController {
	repo := repository.NewAccountsRepository(db, logger)
	useCase := usecases.NewAccountUseCase(repo, logger)
	handler := controllers.NewAccountController(useCase)
	return handler
}
