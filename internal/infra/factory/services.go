package factory

import (
	"github.com/hebertzin/scheduler/internal/core/usecases"
	"github.com/hebertzin/scheduler/internal/domain"
	"github.com/hebertzin/scheduler/internal/infra/db/repository"
	"github.com/hebertzin/scheduler/internal/presentation/controllers"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func ServiceFactory(db *gorm.DB, logger *logrus.Logger) domain.ServicesController {
	repo := repository.NewServicesRepository(db, logger)
	useCase := usecases.NewServicesUseCase(repo, logger)
	handler := controllers.NewServicesController(useCase)
	return handler
}
