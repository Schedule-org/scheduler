package usecases

import (
	"context"

	"github.com/hebertzin/scheduler/internal/core"
	"github.com/hebertzin/scheduler/internal/domains"
	"github.com/hebertzin/scheduler/internal/infra/db/repository"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type AddUserUseCase interface {
	Add(ctx context.Context, payload *domains.User) (*domains.User, *core.Exception)
	FindUserById(ctx context.Context, id string) (*domains.User, *core.Exception)
	FindAllUsers(ctx context.Context) ([]domains.User, *core.Exception)
	FindAllEstablishmentsByUserId(ctx context.Context, user_id string) ([]domains.Establishment, *core.Exception)
}

type AddUserUseCaseImpl struct {
	repo   repository.UserRepository
	logger *logrus.Logger
}

func NewAddUserUseCase(repo repository.UserRepository, logger *logrus.Logger) AddUserUseCase {
	return &AddUserUseCaseImpl{repo: repo, logger: logger}
}

func (uc *AddUserUseCaseImpl) Add(ctx context.Context, payload *domains.User) (*domains.User, *core.Exception) {
	if payload.Name == "" || payload.Email == "" || payload.Password == "" {
		return nil, core.BadRequest(core.WithMessage("Some fields are missing"))
	}

	existentUser, _ := uc.repo.FindUserByEmail(ctx, payload.Email)
	if existentUser != nil {
		return nil, core.Confilct(core.WithMessage("User already exists in the database"))
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, core.Unexpected(core.WithMessage("Error generating password hash"))
	}
	payload.Password = string(hashedPassword)

	user, err := uc.repo.Add(ctx, payload)
	if err != nil {
		return nil, core.Unexpected()
	}

	return user, nil
}

func (ur *AddUserUseCaseImpl) FindUserById(ctx context.Context, id string) (*domains.User, *core.Exception) {
	user, err := ur.repo.FindUserById(ctx, id)
	if err != nil {
		return nil, core.Unexpected(core.WithMessage("error finding user"), core.WithError(err))
	}
	return user, nil
}

func (ur *AddUserUseCaseImpl) FindAllUsers(ctx context.Context) ([]domains.User, *core.Exception) {
	users, err := ur.repo.FindAllUsers(ctx)
	if err != nil {
		return nil, core.Unexpected(core.WithMessage("Some error has been ocurred"))
	}
	return users, nil
}

func (ur *AddUserUseCaseImpl) FindAllEstablishmentsByUserId(ctx context.Context, user_id string) ([]domains.Establishment, *core.Exception) {
	establishments, err := ur.repo.FindAllEstablishmentsByUserId(ctx, user_id)
	if err != nil {
		return nil, core.Unexpected(core.WithMessage("Some error has been ocurred"))
	}
	return establishments, nil
}
