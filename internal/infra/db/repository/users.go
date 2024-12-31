package repository

import (
	"context"

	"github.com/hebertzin/tadix-backend/internal/domains"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UserRepository interface {
	Add(ctx context.Context, category *domains.User) (*domains.User, error)
}

type UserDatabaseRepository struct {
	db     *gorm.DB
	logger *logrus.Logger
}

func NewUserRepository(db *gorm.DB, logger *logrus.Logger) *UserDatabaseRepository {
	return &UserDatabaseRepository{
		db: db,
	}
}

func (repo *UserDatabaseRepository) Add(ctx context.Context, user *domains.User) (*domains.User, error) {
	repo.logger.WithFields(logrus.Fields{
		"method": "Add",
		"user":   user,
	}).Info("Init user creation")

	err := repo.db.WithContext(ctx).Create(user).Error
	if err != nil {
		repo.logger.WithFields(logrus.Fields{
			"method": "Add",
			"error":  err,
		}).Error("Some error has been ocurred trying create a user")
		return nil, err
	}

	repo.logger.WithFields(logrus.Fields{
		"method": "Add",
		"userID": user.Id,
	}).Info("User created sucessfully")

	return user, nil
}
