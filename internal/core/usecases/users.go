package usecases

import (
	"context"

	"github.com/hebertzin/scheduler/internal/core"
	"github.com/hebertzin/scheduler/internal/domains"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type UserUseCase struct {
	repository domains.UserRepository
	logger     *logrus.Logger
}

func NewAddUserUseCase(repository domains.UserRepository, logger *logrus.Logger) domains.UserUseCase {
	return &UserUseCase{repository: repository, logger: logger}
}

func (s *UserUseCase) Add(ctx context.Context, payload *domains.User) (*domains.User, *core.Exception) {
	if payload.Name == "" || payload.Email == "" || payload.Password == "" {
		return nil, core.BadRequest(core.WithMessage("Some fields are missing"))
	}

	existentUser, _ := s.repository.FindUserByEmail(ctx, payload.Email)
	if existentUser != nil {
		return nil, core.Confilct(core.WithMessage("User already exists in the database"))
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

func (s *UserUseCase) FindUserById(ctx context.Context, id string) (*domains.User, *core.Exception) {
	user, err := s.repository.FindUserById(ctx, id)
	if err != nil {
		return nil, core.Unexpected(core.WithMessage("error finding user"), core.WithError(err))
	}
	return user, nil
}

func (s *UserUseCase) FindAllUsers(ctx context.Context) ([]domains.User, *core.Exception) {
	users, err := s.repository.FindAllUsers(ctx)
	if err != nil {
		return nil, core.Unexpected(core.WithMessage("Some error has been ocurred"))
	}
	return users, nil
}

func (s *UserUseCase) FindAllEstablishmentsByUserId(ctx context.Context, user_id string) ([]domains.Establishment, *core.Exception) {
	establishments, err := s.repository.FindAllEstablishmentsByUserId(ctx, user_id)
	if err != nil {
		return nil, core.Unexpected(core.WithMessage("Some error has been ocurred"))
	}
	return establishments, nil
}
