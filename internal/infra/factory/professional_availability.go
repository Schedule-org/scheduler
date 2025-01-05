package factory

import (
	"github.com/hebertzin/scheduler/internal/core/usecases"
	"github.com/hebertzin/scheduler/internal/infra/db/repository"
	"github.com/hebertzin/scheduler/internal/presentation/controllers/v1"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func ProfessionalAvailabilityFactory(db *gorm.DB, logger *logrus.Logger) controllers.ProfessionalAvailabilityController {
	pa := repository.NewProfessionalsAvailabilityRepository(db, logger)
	pc := usecases.NewProfessionalsAvailabilityUseCase(pa, logger)
	ph := controllers.NewProfessionalAvailabilityController(pc)
	return ph
}
