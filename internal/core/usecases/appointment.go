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
	GetAllAppointmentsByProfessionalId(ctx context.Context, professional_id string) ([]domains.Appointment, *core.Exception)
	GetAppointmentById(ctx context.Context, appointment_id string) (*domains.Appointment, *core.Exception)
}

type AppointmentUseCaseImpl struct {
	repo   repository.AppointmentRepository
	logger *logrus.Logger
}

func NewAppointmentUseCase(repo repository.AppointmentRepository, logger *logrus.Logger) AppointmentUseCase {
	return &AppointmentUseCaseImpl{repo: repo, logger: logger}
}

func (ur *AppointmentUseCaseImpl) Add(ctx context.Context, appointment *domains.Appointment) (*domains.Appointment, *core.Exception) {
	appointment, err := ur.repo.Add(ctx, appointment)
	if err != nil {
		return nil, core.Unexpected(core.WithMessage("error creating appointment"), core.WithError(err))
	}
	return appointment, nil
}

func (ur *AppointmentUseCaseImpl) GetAllAppointmentsByProfessionalId(ctx context.Context, professional_id string) ([]domains.Appointment, *core.Exception) {
	appointments, err := ur.repo.GetAllAppointmentsByProfessionalId(ctx, professional_id)
	if err != nil {
		return nil, core.Unexpected(core.WithMessage("error get all appointment by professional id"), core.WithError(err))
	}
	return appointments, nil
}

func (ur *AppointmentUseCaseImpl) GetAppointmentById(ctx context.Context, appointment_id string) (*domains.Appointment, *core.Exception) {
	appointment, err := ur.repo.GetAppointmentById(ctx, appointment_id)
	if err != nil {
		return nil, core.Unexpected(core.WithMessage("error get appointment by id"), core.WithError(err))
	}
	return appointment, nil
}
