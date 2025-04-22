package factory

import (
	"github.com/hebertzin/scheduler/internal/core/usecases"
	"github.com/hebertzin/scheduler/internal/domain"
	"github.com/hebertzin/scheduler/internal/infra/db/repository"
	"github.com/hebertzin/scheduler/internal/presentation/controllers/v1"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func ProfessionalAvailabilityFactory(db *gorm.DB, logger *logrus.Logger) domain.ProfessionalAvailabilityController {
	repo := repository.NewProfessionalsAvailabilityRepository(db, logger)
	useCase := usecases.NewProfessionalsAvailabilityUseCase(repo, logger)
	handler := controllers.NewProfessionalAvailabilityController(useCase)
	return handler
}
