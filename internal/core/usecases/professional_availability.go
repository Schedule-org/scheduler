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
}

type ProfessionalsAvailabilityUseCaseImpl struct {
	repo   repository.ProfessionalsAvailabilityRepository
	logger *logrus.Logger
}

func NewProfessionalsAvailabilityUseCase(repo repository.EstablishmentRepository, logger *logrus.Logger) EstablishmentUseCase {
	return &EstablishmentUserUseCaseImpl{repo: repo, logger: logger}
}

func (uc *ProfessionalsAvailabilityUseCaseImpl) Add(ctx context.Context, availability *domains.ProfessionalAvailability) (*domains.ProfessionalAvailability, *core.Exception) {
	availability, err := uc.repo.Add(ctx, availability)
	if err != nil {
		return nil, core.Unexpected(core.WithMessage("error creating availability"), core.WithError(err))
	}
	return availability, nil
}
