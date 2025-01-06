package usecases

import (
	"context"

	"github.com/hebertzin/scheduler/internal/core"
	"github.com/hebertzin/scheduler/internal/domains"
	"github.com/hebertzin/scheduler/internal/infra/db/repository"
	"github.com/sirupsen/logrus"
)

type ProfessionalsAvailabilityUseCase interface {
	Add(ctx context.Context, availability *domains.ProfessionalAvailability) (*domains.ProfessionalAvailability, *core.Exception)
	GetProfessionalAvailabilityById(ctx context.Context, professional_id string) ([]domains.ProfessionalAvailability, *core.Exception)
}

type ProfessionalsAvailabilityUseCaseImpl struct {
	repo   repository.ProfessionalsAvailabilityRepository
	logger *logrus.Logger
}

func NewProfessionalsAvailabilityUseCase(repo repository.ProfessionalsAvailabilityRepository, logger *logrus.Logger) ProfessionalsAvailabilityUseCase {
	return &ProfessionalsAvailabilityUseCaseImpl{repo: repo, logger: logger}
}

func (uc *ProfessionalsAvailabilityUseCaseImpl) Add(ctx context.Context, availability *domains.ProfessionalAvailability) (*domains.ProfessionalAvailability, *core.Exception) {
	availability, err := uc.repo.Add(ctx, availability)
	if err != nil {
		return nil, core.Unexpected(core.WithMessage("error creating availability"), core.WithError(err))
	}
	return availability, nil
}

func (uc *ProfessionalsAvailabilityUseCaseImpl) GetProfessionalAvailabilityById(ctx context.Context, professionail_id string) ([]domains.ProfessionalAvailability, *core.Exception) {
	availability, err := uc.repo.GetProfessionalAvailabilityById(ctx, professionail_id)
	if err != nil {
		return nil, core.Unexpected(core.WithMessage("error get professional availability"), core.WithError(err))
	}
	return availability, nil
}
