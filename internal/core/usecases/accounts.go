package usecases

import (
	"context"

	"github.com/hebertzin/scheduler/internal/core"
	"github.com/hebertzin/scheduler/internal/domain"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type AccountUseCase struct {
	repository domain.AccountRepository
	logger     *logrus.Logger
}

func NewAccountUseCase(repository domain.AccountRepository, logger *logrus.Logger) domain.AccountUseCase {
	return &AccountUseCase{repository: repository, logger: logger}
}

func (s *AccountUseCase) Add(ctx context.Context, payload *domain.Account) (*domain.Account, *core.Exception) {
	if payload.Name == "" || payload.Email == "" || payload.Password == "" {
		return nil, core.BadRequest(core.WithMessage("Some fields are missing"))
	}

	existentUser, _ := s.repository.FindAccountByEmail(ctx, payload.Email)
	if existentUser != nil {
		return nil, core.Confilct(core.WithMessage("Account already exists in the database"))
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, core.Unexpected(core.WithMessage("Error generating password hash"))
	}
	payload.Password = string(hashedPassword)

	user, err := s.repository.Add(ctx, payload)
	if err != nil {
		return nil, core.Unexpected()
	}

	return user, nil
}

func (s *AccountUseCase) FindAccountById(ctx context.Context, id string) (*domain.Account, *core.Exception) {
	user, err := s.repository.FindAccountById(ctx, id)
	if err != nil {
		return nil, core.Unexpected(core.WithMessage("Error finding user"), core.WithError(err))
	}
	return user, nil
}

func (s *AccountUseCase) FindAllAccounts(ctx context.Context) ([]domain.Account, *core.Exception) {
	users, err := s.repository.FindAllAccounts(ctx)
	if err != nil {
		return nil, core.Unexpected(core.WithMessage("Some error has been ocurred"))
	}
	return users, nil
}

func (s *AccountUseCase) FindAllEstablishmentsByAccountId(ctx context.Context, user_id string) ([]domain.Establishment, *core.Exception) {
	establishments, err := s.repository.FindAllEstablishmentsByAccountId(ctx, user_id)
	if err != nil {
		return nil, core.Unexpected(core.WithMessage("Some error has been ocurred"))
	}
	return establishments, nil
}
