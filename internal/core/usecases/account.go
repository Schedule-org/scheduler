package usecases

import (
	"context"

	"github.com/hebertzin/scheduler/internal/core"
	"github.com/hebertzin/scheduler/internal/domain"
	"github.com/sirupsen/logrus"
)

type ClientUseCase struct {
	repository domain.ClientUseCase
	logger     *logrus.Logger
}

func NewClientUseCase(repository domain.ClientRepository, logger *logrus.Logger) domain.ClientUseCase {
	return &ClientUseCase{repository: repository, logger: logger}
}

func (s *ClientUseCase) Add(ctx context.Context, account *domain.Client) (*domain.Client, *core.Exception) {
	account, err := s.repository.Add(ctx, account)
	if err != nil {
		return nil, core.Unexpected(core.WithMessage("Error creating account"), core.WithError(err))
	}
	return account, nil
}
