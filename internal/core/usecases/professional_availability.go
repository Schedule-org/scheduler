package usecases

import (
	"context"

	"github.com/hebertzin/scheduler/internal/core"
	"github.com/hebertzin/scheduler/internal/domains"
	"github.com/hebertzin/scheduler/internal/infra/db/repository"
	"github.com/sirupsen/logrus"
)

type ProfessionalsAvailabilityUseCase struct {
	repository repository.ProfessionalsAvailabilityRepository
	logger     *logrus.Logger
}

func NewProfessionalsAvailabilityUseCase(repository repository.ProfessionalsAvailabilityRepository, logger *logrus.Logger) domains.ProfessionalsAvailabilityUseCase {
	return &ProfessionalsAvailabilityUseCase{repository: repository, logger: logger}
}

func (s *ProfessionalsAvailabilityUseCase) Add(ctx context.Context, availability *domains.ProfessionalAvailability) (*domains.ProfessionalAvailability, *core.Exception) {
	availability, err := s.repository.Add(ctx, availability)
	if err != nil {
		return nil, core.Unexpected(core.WithMessage("error creating availability"), core.WithError(err))
	}
	return availability, nil
}

func (s *ProfessionalsAvailabilityUseCase) GetProfessionalAvailabilityById(ctx context.Context, professionail_id string) ([]domains.ProfessionalAvailability, *core.Exception) {
	availability, err := s.repository.GetProfessionalAvailabilityById(ctx, professionail_id)
	if err != nil {
		return nil, core.Unexpected(core.WithMessage("error get professional availability"), core.WithError(err))
	}
	return availability, nil
}
