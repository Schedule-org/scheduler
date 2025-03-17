package usecases

import (
	"context"

	"github.com/hebertzin/scheduler/internal/core"
	"github.com/hebertzin/scheduler/internal/domains"
	"github.com/hebertzin/scheduler/internal/infra/db/repository"
	"github.com/sirupsen/logrus"
)

type AppointmentUseCase struct {
	repository repository.AppointmentRepository
	logger     *logrus.Logger
}

func NewAppointmentUseCase(repository repository.AppointmentRepository, logger *logrus.Logger) domains.AppointmentUseCase {
	return &AppointmentUseCase{repository: repository, logger: logger}
}

func (s *AppointmentUseCase) Add(ctx context.Context, appointment *domains.Appointment) (*domains.Appointment, *core.Exception) {
	appointment, err := s.repository.Add(ctx, appointment)
	if err != nil {
		return nil, core.Unexpected(core.WithMessage("error creating appointment"), core.WithError(err))
	}
	return appointment, nil
}

func (s *AppointmentUseCase) GetAllAppointmentsByProfessionalId(ctx context.Context, professional_id string) ([]domains.Appointment, *core.Exception) {
	appointments, err := s.repository.GetAllAppointmentsByProfessionalId(ctx, professional_id)
	if err != nil {
		return nil, core.Unexpected(core.WithMessage("error get all appointment by professional id"), core.WithError(err))
	}
	return appointments, nil
}

func (s *AppointmentUseCase) GetAppointmentById(ctx context.Context, appointment_id string) (*domains.Appointment, *core.Exception) {
	appointment, err := s.repository.GetAppointmentById(ctx, appointment_id)
	if err != nil {
		return nil, core.Unexpected(core.WithMessage("error get appointment by id"), core.WithError(err))
	}
	return appointment, nil
}
