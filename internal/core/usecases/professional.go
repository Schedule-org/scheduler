package usecases

import (
	"context"

	"github.com/hebertzin/scheduler/internal/core"
	"github.com/hebertzin/scheduler/internal/domains"
	"github.com/hebertzin/scheduler/internal/infra/db/repository"
	"github.com/sirupsen/logrus"
)

type ProfessionalsUseCase struct {
	repo   repository.ProfessionalsRepository
	logger *logrus.Logger
}

func NewProfissionalUseCase(repo repository.ProfessionalsRepository, logger *logrus.Logger) domains.ProfessionalsUseCase {
	return &ProfessionalsUseCase{repo: repo, logger: logger}
}

func (uc *ProfessionalsUseCase) FindProfessionalById(ctx context.Context, id string) (*domains.Professionals, *core.Exception) {
	professional, err := uc.repo.FindProfessionalById(ctx, id)
	if err != nil {
		return nil, core.Unexpected(core.WithMessage("error finding professional"), core.WithError(err))
	}
	return professional, nil
}

func (uc *ProfessionalsUseCase) Add(ctx context.Context, payload *domains.Professionals) (*domains.Professionals, *core.Exception) {
	if payload.Name == "" || payload.Role == "" || payload.EstablishmentId == "" {
		return nil, core.BadRequest(core.WithMessage("Some fields are missing"))
	}
	professional, err := uc.repo.Add(ctx, payload)
	if err != nil {
		return nil, core.Unexpected()
	}
	return professional, nil
}

func (uc *ProfessionalsUseCase) UpdateProfessionalById(ctx context.Context, professionail_id string, professionalData *domains.Professionals) (*domains.Professionals, *core.Exception) {
	professional, err := uc.repo.UpdateProfessionalById(ctx, professionail_id, professionalData)
	if err != nil {
		return nil, core.Unexpected()
	}
	return professional, nil
}
