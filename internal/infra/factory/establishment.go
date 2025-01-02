package factory

import (
	"github.com/hebertzin/scheduler/internal/core/usecases"
	"github.com/hebertzin/scheduler/internal/infra/db/repository"
	"github.com/hebertzin/scheduler/internal/presentation/controllers/v1"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func EstablishmentFactory(db *gorm.DB, logger *logrus.Logger) controllers.EstablishmentController {
	er := repository.NewEstablishmentRepository(db, logger)
	ec := usecases.NewEstablishmentUseCase(er, logger)
	eh := controllers.NewEstablishmentController(ec)
	return eh
}
