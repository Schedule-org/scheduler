package factory

import (
	"github.com/hebertzin/scheduler/internal/core/usecases"
	"github.com/hebertzin/scheduler/internal/infra/db/repository"
	"github.com/hebertzin/scheduler/internal/presentation/controllers/v1"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func ServiceFactory(db *gorm.DB, logger *logrus.Logger) controllers.ServicesController {
	sr := repository.NewServicesRepository(db, logger)
	sc := usecases.NewServicesUseCase(sr, logger)
	sh := controllers.NewServicesController(sc)
	return sh
}
