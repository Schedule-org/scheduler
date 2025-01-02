package factory

import (
	"github.com/hebertzin/scheduler/internal/core/usecases"
	"github.com/hebertzin/scheduler/internal/infra/db/repository"
	"github.com/hebertzin/scheduler/internal/presentation/controllers/v1"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func ProfessionalFactory(db *gorm.DB, logger *logrus.Logger) controllers.ProfessionalsController {
	pr := repository.NewProfessionalsRepository(db, logger)
	pc := usecases.NewProfissionalUseCase(pr, logger)
	ph := controllers.NewProfessionalController(pc)
	return ph
}
