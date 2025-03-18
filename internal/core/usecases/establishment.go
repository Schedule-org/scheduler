package usecases

import (
	"context"

	"github.com/hebertzin/scheduler/internal/core"
	"github.com/hebertzin/scheduler/internal/domains"
	"github.com/sirupsen/logrus"
)

type EstablishmentUserUseCase struct {
	repository domains.EstablishmentRepository
	logger     *logrus.Logger
}

func NewEstablishmentUseCase(repository domains.EstablishmentRepository, logger *logrus.Logger) domains.EstablishmentUseCase {
	return &EstablishmentUserUseCase{repository: repository, logger: logger}
}

func (s *EstablishmentUserUseCase) FindEstablishmentById(ctx context.Context, id string) (*domains.Establishment, *core.Exception) {
	establishment, err := s.repository.FindEstablishmentById(ctx, id)
	if err != nil {
		return nil, core.Unexpected(core.WithMessage("error finding establishment"), core.WithError(err))
	}
	return establishment, nil
}

func (s *EstablishmentUserUseCase) Add(ctx context.Context, payload *domains.Establishment) (*domains.Establishment, *core.Exception) {
	establishment, err := s.repository.Add(ctx, payload)
	if err != nil {
		return nil, core.Unexpected()
	}
	return establishment, nil
}

func (s *EstablishmentUserUseCase) GetAllProfessionalsByEstablishmentId(ctx context.Context, establishment_id string) ([]domains.Professionals, *core.Exception) {
	professionails, err := s.repository.GetAllProfessionalsByEstablishmentId(ctx, establishment_id)
	if err != nil {
		return nil, core.Unexpected()
	}
	return professionails, nil
}

func (s *EstablishmentUserUseCase) UpdateEstablishmentById(ctx context.Context, establishment_id string, establishmentData *domains.Establishment) (*domains.Establishment, *core.Exception) {
	establishment, err := s.repository.UpdateEstablishmentById(ctx, establishment_id, establishmentData)
	if err != nil {
		return nil, core.Unexpected(core.WithMessage("Some error has been ocurred trying update a establishment"))
	}
	return establishment, nil
}

func (s *EstablishmentUserUseCase) GetEstablishmentReport(ctx context.Context, establishment_id string) (*domains.EstablishmentReport, *core.Exception) {
	stats, err := s.repository.GetEstablishmentReport(ctx, establishment_id)
	if err != nil {
		return nil, core.Unexpected(core.WithMessage("Some error has been ocurred trying update a establishment"))
	}
	return stats, nil
}
