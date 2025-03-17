package usecases

import (
	"context"

	"github.com/hebertzin/scheduler/internal/core"
	"github.com/hebertzin/scheduler/internal/domains"
	"github.com/hebertzin/scheduler/internal/infra/db/repository"
	"github.com/sirupsen/logrus"
)

type ServicesUseCase struct {
	repository repository.ServicesRepository
	logger     *logrus.Logger
}

func NewServicesUseCase(repository repository.ServicesRepository, logger *logrus.Logger) domains.ServicesUseCase {
	return &ServicesUseCase{repository: repository, logger: logger}
}

func (s *ServicesUseCase) FindServiceById(ctx context.Context, id string) (*domains.Services, *core.Exception) {
	service, err := s.repository.FindServiceById(ctx, id)
	if err != nil {
		return nil, core.Unexpected(core.WithMessage("error finding service"), core.WithError(err))
	}
	return service, nil
}

func (s *ServicesUseCase) Add(ctx context.Context, payload *domains.Services) (*domains.Services, *core.Exception) {
	if payload.Name == "" || payload.Duration == "" {
		return nil, core.BadRequest(core.WithMessage("Some fields are missing"))
	}
	service, err := s.repository.Add(ctx, payload)
	if err != nil {
		return nil, core.Unexpected()
	}
	return service, nil
}

func (s *ServicesUseCase) GetAllServicesByProfessionalId(ctx context.Context, professional_id string) ([]domains.Services, *core.Exception) {
	services, err := s.repository.GetAllServicesByProfessionalId(ctx, professional_id)
	if err != nil {
		return nil, core.Unexpected()
	}
	return services, nil
}
