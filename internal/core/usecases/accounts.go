package usecases

import (
	"context"
	"regexp"

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
	isValidEmail := validateAccountEmail(payload.Email)
	if !isValidEmail {
		return nil, core.BadRequest(core.WithMessage("Email invalid"))
	}

	account, _ := s.repository.FindAccountByEmail(ctx, payload.Email)
	if account == nil {
		return nil, core.Confilct(core.WithMessage("Account already exists in the database"))
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, core.Unexpected(core.WithMessage("Error generating password hash"))
	}
	payload.Password = string(hash)

	a, err := s.repository.Add(ctx, payload)
	if err != nil {
		return nil, core.Unexpected()
	}

	return a, nil
}

func (s *AccountUseCase) FindAccountById(ctx context.Context, id string) (*domain.Account, *core.Exception) {
	account, err := s.repository.FindAccountById(ctx, id)
	if err != nil {
		return nil, core.Unexpected(core.WithMessage("Error finding account"), core.WithError(err))
	}
	return account, nil
}

func (s *AccountUseCase) FindAllAccounts(ctx context.Context) ([]domain.Account, *core.Exception) {
	account, err := s.repository.FindAllAccounts(ctx)
	if err != nil {
		return nil, core.Unexpected(core.WithMessage("Some error has been ocurred"))
	}
	return account, nil
}

func (s *AccountUseCase) FindAllEstablishmentsByAccountId(ctx context.Context, account_id string) ([]domain.Establishment, *core.Exception) {
	establishments, err := s.repository.FindAllEstablishmentsByAccountId(ctx, account_id)
	if err != nil {
		return nil, core.Unexpected(core.WithMessage("Some error has been ocurred"))
	}
	return establishments, nil
}

func validateAccountEmail(email string) bool {
	regex := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	return regex.MatchString(email)
}
