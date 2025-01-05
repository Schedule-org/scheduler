package usecases

import (
	"context"

	"github.com/hebertzin/scheduler/internal/core"
	"github.com/hebertzin/scheduler/internal/domains"
	"github.com/hebertzin/scheduler/internal/infra/db/repository"
	"github.com/sirupsen/logrus"
)

type ProfessionalsUseCase interface {
	Add(ctx context.Context, payload *domains.Professionals) (*domains.Professionals, *core.Exception)
	FindProfessionalById(ctx context.Context, id string) (*domains.Professionals, *core.Exception)
}

type ProfessionalsUseCaseImpl struct {
	repo   repository.ProfessionalsRepository
	logger *logrus.Logger
}

func NewProfissionalUseCase(repo repository.ProfessionalsRepository, logger *logrus.Logger) ProfessionalsUseCase {
	return &ProfessionalsUseCaseImpl{repo: repo, logger: logger}
}

func (uc *ProfessionalsUseCaseImpl) FindProfessionalById(ctx context.Context, id string) (*domains.Professionals, *core.Exception) {
	professional, err := uc.repo.FindProfessionalById(ctx, id)
	if err != nil {
		return nil, core.Unexpected(core.WithMessage("error finding professional"), core.WithError(err))
	}
	return professional, nil
}

func (uc *ProfessionalsUseCaseImpl) Add(ctx context.Context, payload *domains.Professionals) (*domains.Professionals, *core.Exception) {
	if payload.Name == "" || payload.Role == "" || payload.EstablishmentId == "" {
		return nil, core.BadRequest(core.WithMessage("Some fields are missing"))
	}
	professional, err := uc.repo.Add(ctx, payload)
	if err != nil {
		return nil, core.Unexpected()
	}
	return professional, nil
}
