package factory

import (
	"github.com/hebertzin/scheduler/internal/core/usecases"
	"github.com/hebertzin/scheduler/internal/domains"
	"github.com/hebertzin/scheduler/internal/infra/db/repository"
	"github.com/hebertzin/scheduler/internal/presentation/controllers/v1"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func EstablishmentFactory(db *gorm.DB, logger *logrus.Logger) domains.EstablishmentController {
	repo := repository.NewEstablishmentRepository(db, logger)
	useCase := usecases.NewEstablishmentUseCase(repo, logger)
	handler := controllers.NewEstablishmentController(useCase)
	return handler
}
