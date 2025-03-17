package usecases

import (
	"context"

	"github.com/hebertzin/scheduler/internal/core"
	"github.com/hebertzin/scheduler/internal/domains"
	"github.com/hebertzin/scheduler/internal/infra/db/repository"
	"github.com/sirupsen/logrus"
)

type ServicesUseCase struct {
	repo   repository.ServicesRepository
	logger *logrus.Logger
}

func NewServicesUseCase(repo repository.ServicesRepository, logger *logrus.Logger) domains.ServicesUseCase {
	return &ServicesUseCase{repo: repo, logger: logger}
}

func (ur *ServicesUseCase) FindServiceById(ctx context.Context, id string) (*domains.Services, *core.Exception) {
	service, err := ur.repo.FindServiceById(ctx, id)
	if err != nil {
		return nil, core.Unexpected(core.WithMessage("error finding service"), core.WithError(err))
	}
	return service, nil
}

func (uc *ServicesUseCase) Add(ctx context.Context, payload *domains.Services) (*domains.Services, *core.Exception) {
	if payload.Name == "" || payload.Duration == "" {
		return nil, core.BadRequest(core.WithMessage("Some fields are missing"))
	}
	service, err := uc.repo.Add(ctx, payload)
	if err != nil {
		return nil, core.Unexpected()
	}
	return service, nil
}

func (uc *ServicesUseCase) GetAllServicesByProfessionalId(ctx context.Context, professional_id string) ([]domains.Services, *core.Exception) {
	services, err := uc.repo.GetAllServicesByProfessionalId(ctx, professional_id)
	if err != nil {
		return nil, core.Unexpected()
	}
	return services, nil
}
