package usecases

import (
	"context"

	"github.com/hebertzin/scheduler/internal/core"
	"github.com/hebertzin/scheduler/internal/domains"
	"github.com/hebertzin/scheduler/internal/infra/db/repository"
	"github.com/sirupsen/logrus"
)

type EstablishmentUseCase interface {
	Add(ctx context.Context, payload *domains.Establishment) (*domains.Establishment, *core.Exception)
	GetAllProfessionalsByEstablishmentId(ctx context.Context, establishment_id string) ([]domains.Professionals, *core.Exception)
	FindEstablishmentById(ctx context.Context, establishment_id string) (*domains.Establishment, *core.Exception)
	UpdateEstablishmentById(ctx context.Context, establishment_id string, establishmentData *domains.Establishment) (*domains.Establishment, *core.Exception)
}

type EstablishmentUserUseCaseImpl struct {
	repo   repository.EstablishmentRepository
	logger *logrus.Logger
}

func NewEstablishmentUseCase(repo repository.EstablishmentRepository, logger *logrus.Logger) EstablishmentUseCase {
	return &EstablishmentUserUseCaseImpl{repo: repo, logger: logger}
}

func (ur *EstablishmentUserUseCaseImpl) FindEstablishmentById(ctx context.Context, id string) (*domains.Establishment, *core.Exception) {
	establishment, err := ur.repo.FindEstablishmentById(ctx, id)
	if err != nil {
		return nil, core.Unexpected(core.WithMessage("error finding establishment"), core.WithError(err))
	}
	return establishment, nil
}

func (uc *EstablishmentUserUseCaseImpl) Add(ctx context.Context, payload *domains.Establishment) (*domains.Establishment, *core.Exception) {
	if payload.Name == "" || payload.City == "" || payload.PostalCode == "" || payload.State == "" {
		return nil, core.BadRequest(core.WithMessage("Some fields are missing"))
	}
	establishment, err := uc.repo.Add(ctx, payload)
	if err != nil {
		return nil, core.Unexpected()
	}

	uc.logger.WithFields(logrus.Fields{
		"method":        "Add",
		"establishment": establishment.Name,
	}).Info("establishment created successfully")

	return establishment, nil
}

func (uc *EstablishmentUserUseCaseImpl) GetAllProfessionalsByEstablishmentId(ctx context.Context, establishment_id string) ([]domains.Professionals, *core.Exception) {
	professionails, err := uc.repo.GetAllProfessionalsByEstablishmentId(ctx, establishment_id)
	if err != nil {
		return nil, core.Unexpected()
	}
	return professionails, nil
}

func (uc *EstablishmentUserUseCaseImpl) UpdateEstablishmentById(ctx context.Context, establishment_id string, establishmentData *domains.Establishment) (*domains.Establishment, *core.Exception) {
	establishment, err := uc.repo.UpdateEstablishmentById(ctx, establishment_id, establishmentData)
	if err != nil {
		return nil, core.Unexpected(core.WithMessage("Some error has been ocurred trying update a establishment"))
	}
	return establishment, nil
}
