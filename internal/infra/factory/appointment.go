package factory

import (
	"github.com/hebertzin/scheduler/internal/core/usecases"
	"github.com/hebertzin/scheduler/internal/domains"
	"github.com/hebertzin/scheduler/internal/infra/db/repository"
	"github.com/hebertzin/scheduler/internal/presentation/controllers/v1"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func AppointmentFactory(db *gorm.DB, logger *logrus.Logger) domains.AppointmentController {
	repo := repository.NewAppointmentRepository(db, logger)
	useCase := usecases.NewAppointmentUseCase(repo, logger)
	handler := controllers.NewAppointmentController(useCase)
	return handler
}
