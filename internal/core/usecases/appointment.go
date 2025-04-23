package usecases

import (
	"context"

	"github.com/hebertzin/scheduler/internal/core"
	"github.com/hebertzin/scheduler/internal/domain"
	"github.com/sirupsen/logrus"
)

type AppointmentUseCase struct {
	repository domain.AppointmentRepository
	logger     *logrus.Logger
}

func NewAppointmentUseCase(repository domain.AppointmentRepository, logger *logrus.Logger) domain.AppointmentUseCase {
	return &AppointmentUseCase{repository: repository, logger: logger}
}

func (s *AppointmentUseCase) Add(ctx context.Context, appointment *domain.Appointment) (*domain.Appointment, *core.Exception) {
	appointment, err := s.repository.Add(ctx, appointment)
	if err != nil {
		return nil, core.Unexpected(core.WithMessage("Error creating appointment"), core.WithError(err))
	}
	return appointment, nil
}

func (s *AppointmentUseCase) GetAllAppointmentsByProfessionalId(ctx context.Context, professional_id string) ([]domain.Appointment, *core.Exception) {
	appointments, err := s.repository.GetAllAppointmentsByProfessionalId(ctx, professional_id)
	if err != nil {
		return nil, core.Unexpected(core.WithMessage("Error get all appointment by professional id"), core.WithError(err))
	}
	return appointments, nil
}

func (s *AppointmentUseCase) GetAppointmentById(ctx context.Context, appointment_id string) (*domain.Appointment, *core.Exception) {
	appointment, err := s.repository.GetAppointmentById(ctx, appointment_id)
	if err != nil {
		return nil, core.Unexpected(core.WithMessage("Error get appointment by id"), core.WithError(err))
	}
	return appointment, nil
}

func (s *AppointmentUseCase) DeleteAppointment(ctx context.Context, appointment_id string) *core.Exception {
	err := s.repository.DeleteAppointment(ctx, appointment_id)
	if err != nil {
		return core.Unexpected(core.WithMessage("Error get appointment by id"), core.WithError(err))
	}
	return nil
}
