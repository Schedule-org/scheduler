package factory

import (
	"github.com/hebertzin/scheduler/internal/core/usecases"
	"github.com/hebertzin/scheduler/internal/domain"
	"github.com/hebertzin/scheduler/internal/infra/db/repository"
	"github.com/hebertzin/scheduler/internal/presentation/controllers"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func ProfessionalFactory(db *gorm.DB, logger *logrus.Logger) domain.ProfessionalsController {
	repo := repository.NewProfessionalsRepository(db, logger)
	useCase := usecases.NewProfissionalUseCase(repo, logger)
	handler := controllers.NewProfessionalController(useCase)
	return handler
}
