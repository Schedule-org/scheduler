package factory

import (
	"github.com/hebertzin/scheduler/internal/core/usecases"
	"github.com/hebertzin/scheduler/internal/domain"
	"github.com/hebertzin/scheduler/internal/infra/db/repository"
	"github.com/hebertzin/scheduler/internal/presentation/controllers"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func EstablishmentFactory(db *gorm.DB, logger *logrus.Logger) domain.EstablishmentController {
	repo := repository.NewEstablishmentRepository(db, logger)
	useCase := usecases.NewEstablishmentUseCase(repo, logger)
	handler := controllers.NewEstablishmentController(useCase)
	return handler
}
