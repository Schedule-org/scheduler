package usecases

import (
	"context"

	"github.com/hebertzin/scheduler/internal/core"
	"github.com/hebertzin/scheduler/internal/domain"
	"github.com/sirupsen/logrus"
)

type ProfessionalsAvailabilityUseCase struct {
	repository domain.ProfessionalsAvailabilityRepository
	logger     *logrus.Logger
}

func NewProfessionalsAvailabilityUseCase(repository domain.ProfessionalsAvailabilityRepository, logger *logrus.Logger) domain.ProfessionalsAvailabilityUseCase {
	return &ProfessionalsAvailabilityUseCase{repository: repository, logger: logger}
}

func (s *ProfessionalsAvailabilityUseCase) Add(ctx context.Context, availability *domain.ProfessionalAvailability) (*domain.ProfessionalAvailability, *core.Exception) {
	availability, err := s.repository.Add(ctx, availability)
	if err != nil {
		return nil, core.Unexpected(core.WithMessage("Error creating availability"), core.WithError(err))
	}
	return availability, nil
}

func (s *ProfessionalsAvailabilityUseCase) GetProfessionalAvailabilityById(ctx context.Context, professionail_id string) ([]domain.ProfessionalAvailability, *core.Exception) {
	availability, err := s.repository.GetProfessionalAvailabilityById(ctx, professionail_id)
	if err != nil {
		return nil, core.Unexpected(core.WithMessage("Error get professional availability"), core.WithError(err))
	}
	return availability, nil
}
