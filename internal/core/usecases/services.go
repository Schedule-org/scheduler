package usecases

import (
	"context"

	"github.com/hebertzin/scheduler/internal/core"
	"github.com/hebertzin/scheduler/internal/domain"
	"github.com/sirupsen/logrus"
)

type ServicesUseCase struct {
	repository domain.ServicesRepository
	logger     *logrus.Logger
}

func NewServicesUseCase(repository domain.ServicesRepository, logger *logrus.Logger) domain.ServicesUseCase {
	return &ServicesUseCase{repository: repository, logger: logger}
}

func (s *ServicesUseCase) FindServiceById(ctx context.Context, id string) (*domain.Services, *core.Exception) {
	service, err := s.repository.FindServiceById(ctx, id)
	if err != nil {
		return nil, core.Unexpected(core.WithMessage("Error finding service"), core.WithError(err))
	}
	return service, nil
}

func (s *ServicesUseCase) Add(ctx context.Context, payload *domain.Services) (*domain.Services, *core.Exception) {
	if payload.Name == "" || payload.Duration == "" {
		return nil, core.BadRequest(core.WithMessage("Some fields are missing"))
	}
	service, err := s.repository.Add(ctx, payload)
	if err != nil {
		return nil, core.Unexpected()
	}
	return service, nil
}

func (s *ServicesUseCase) GetAllServicesByProfessionalId(ctx context.Context, professional_id string) ([]domain.Services, *core.Exception) {
	services, err := s.repository.GetAllServicesByProfessionalId(ctx, professional_id)
	if err != nil {
		return nil, core.Unexpected()
	}
	return services, nil
}
