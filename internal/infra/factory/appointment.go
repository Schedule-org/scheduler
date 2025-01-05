package factory

import (
	"github.com/hebertzin/scheduler/internal/core/usecases"
	"github.com/hebertzin/scheduler/internal/infra/db/repository"
	"github.com/hebertzin/scheduler/internal/presentation/controllers/v1"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func AppointmentFactory(db *gorm.DB, logger *logrus.Logger) controllers.AppointmentController {
	er := repository.NewAppointmentRepository(db, logger)
	ec := usecases.NewAppointmentUseCase(er, logger)
	eh := controllers.NewAppointmentController(ec)
	return eh
}
