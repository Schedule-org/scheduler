package usecases

import (
	"context"

	"github.com/hebertzin/scheduler/internal/core"
	"github.com/hebertzin/scheduler/internal/domain"
	"github.com/sirupsen/logrus"
)

type AccountUseCase struct {
	repository domain.AccountUseCase
	logger     *logrus.Logger
}

func NewAccountUseCase(repository domain.AccountRepository, logger *logrus.Logger) domain.AccountUseCase {
	return &AccountUseCase{repository: repository, logger: logger}
}

func (s *AccountUseCase) Add(ctx context.Context, account *domain.Account) (*domain.Account, *core.Exception) {
	account, err := s.repository.Add(ctx, account)
	if err != nil {
		return nil, core.Unexpected(core.WithMessage("Error creating account"), core.WithError(err))
	}
	return account, nil
}
