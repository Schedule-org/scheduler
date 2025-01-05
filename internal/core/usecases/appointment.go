package usecases

import (
	"context"

	"github.com/hebertzin/scheduler/internal/core"
	"github.com/hebertzin/scheduler/internal/domains"
	"github.com/hebertzin/scheduler/internal/infra/db/repository"
	"github.com/sirupsen/logrus"
)

type AppointmentUseCase interface {
	Add(ctx context.Context, appointment *domains.Appointment) (*domains.Appointment, *core.Exception)
}

type AppointmentUseCaseImpl struct {
	repo   repository.AppointmentRepository
	logger *logrus.Logger
}

func NewAppointmentUseCase(repo repository.AppointmentRepository, logger *logrus.Logger) AppointmentUseCase {
	return &AppointmentUseCaseImpl{repo: repo, logger: logger}
}

func (ur *AppointmentUseCaseImpl) Add(ctx context.Context, appointment *domains.Appointment) (*domains.Appointment, *core.Exception) {
	service, err := ur.repo.Add(ctx, appointment)
	if err != nil {
		return nil, core.Unexpected(core.WithMessage("error creating appointment"), core.WithError(err))
	}
	return service, nil
}
