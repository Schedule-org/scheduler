package usecases

import (
	"context"

	"github.com/hebertzin/scheduler/internal/core"
	"github.com/hebertzin/scheduler/internal/domains"
	"github.com/hebertzin/scheduler/internal/infra/db/repository"
	"github.com/sirupsen/logrus"
)

type EstablishmentUserUseCase struct {
	repo   repository.EstablishmentRepository
	logger *logrus.Logger
}

func NewEstablishmentUseCase(repo repository.EstablishmentRepository, logger *logrus.Logger) domains.EstablishmentUseCase {
	return &EstablishmentUserUseCase{repo: repo, logger: logger}
}

func (uc *EstablishmentUserUseCase) FindEstablishmentById(ctx context.Context, id string) (*domains.Establishment, *core.Exception) {
	establishment, err := uc.repo.FindEstablishmentById(ctx, id)
	if err != nil {
		return nil, core.Unexpected(core.WithMessage("error finding establishment"), core.WithError(err))
	}
	return establishment, nil
}

func (uc *EstablishmentUserUseCase) Add(ctx context.Context, payload *domains.Establishment) (*domains.Establishment, *core.Exception) {
	if payload.Name == "" || payload.City == "" || payload.PostalCode == "" || payload.State == "" {
		return nil, core.BadRequest(core.WithMessage("Some fields are missing"))
	}
	establishment, err := uc.repo.Add(ctx, payload)
	if err != nil {
		return nil, core.Unexpected()
	}
	return establishment, nil
}

func (uc *EstablishmentUserUseCase) GetAllProfessionalsByEstablishmentId(ctx context.Context, establishment_id string) ([]domains.Professionals, *core.Exception) {
	professionails, err := uc.repo.GetAllProfessionalsByEstablishmentId(ctx, establishment_id)
	if err != nil {
		return nil, core.Unexpected()
	}
	return professionails, nil
}

func (uc *EstablishmentUserUseCase) UpdateEstablishmentById(ctx context.Context, establishment_id string, establishmentData *domains.Establishment) (*domains.Establishment, *core.Exception) {
	establishment, err := uc.repo.UpdateEstablishmentById(ctx, establishment_id, establishmentData)
	if err != nil {
		return nil, core.Unexpected(core.WithMessage("Some error has been ocurred trying update a establishment"))
	}
	return establishment, nil
}

func (uc *EstablishmentUserUseCase) GetEstablishmentReport(ctx context.Context, establishment_id string) (*domains.EstablishmentReport, *core.Exception) {
	stats, err := uc.repo.GetEstablishmentReport(ctx, establishment_id)
	if err != nil {
		return nil, core.Unexpected(core.WithMessage("Some error has been ocurred trying update a establishment"))
	}
	return stats, nil
}
