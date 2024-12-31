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

	_, err := uc.repo.FindUserByEmail(ctx, payload.Email)
	if err == nil {
		uc.logger.WithFields(logrus.Fields{
			"method": "Add",
			"error":  err,
		}).Error("User already exists")
		return nil, core.Confilct(core.WithMessage("User already exist in database"))
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	if err != nil {
		uc.logger.WithFields(logrus.Fields{
			"method": "Add",
			"error":  err,
		}).Error("Error generating password hash")
		return nil, core.Unexpected(core.WithMessage("error generating password hash"))
	}
	payload.Password = string(hashedPassword)

	user, err := uc.repo.Add(ctx, payload)
	if err != nil {
		uc.logger.WithFields(logrus.Fields{
			"method": "Add",
			"error":  err,
		}).Error("Error creating user")
		return nil, core.Unexpected()
	}

	uc.logger.WithFields(logrus.Fields{
		"method": "Add",
		"userID": user.Id,
	}).Info("User created successfully")

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
	result, err := ur.repo.FindAllUsers(ctx)
	if err != nil {
		return nil, core.Unexpected(core.WithMessage("Some error has been ocurred"))
	}
	return result, nil
}
