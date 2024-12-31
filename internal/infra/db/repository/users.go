package repository

import (
	"context"
	"errors"

	"github.com/hebertzin/tadix-backend/internal/domains"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// Interface UserRepository define os métodos para operações no banco de dados relacionadas a usuários.
type UserRepository interface {
	Add(ctx context.Context, user *domains.User) (*domains.User, error)
	FindUserByEmail(ctx context.Context, email string) (*domains.User, error)
	FindUserById(ctx context.Context, id string) (*domains.User, error)
	FindAllUsers(ctx context.Context) ([]domains.User, error)
}

// Estrutura UserDatabaseRepository implementa UserRepository e opera sobre um banco de dados GORM.
type UserDatabaseRepository struct {
	db     *gorm.DB
	logger *logrus.Logger
}

func NewUserRepository(db *gorm.DB, logger *logrus.Logger) *UserDatabaseRepository {
	return &UserDatabaseRepository{
		db:     db,
		logger: logger,
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
		}).Error("Error occurred while creating user")
		return nil, err
	}

	repo.logger.WithFields(logrus.Fields{
		"method": "Add",
		"userID": user.Id,
	}).Info("User created successfully")

	return user, nil
}

func (repo *UserDatabaseRepository) FindUserById(ctx context.Context, id string) (*domains.User, error) {
	var user domains.User
	err := repo.db.WithContext(ctx).Where("id = ?", id).First(&user).Error
	if err != nil {
		repo.logger.WithFields(logrus.Fields{
			"method": "FindUserById",
			"userID": id,
			"error":  err,
		}).Error("Error finding user by ID")

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	repo.logger.WithFields(logrus.Fields{
		"method": "FindUserById",
		"userID": id,
	}).Info("User found successfully")

	return &user, nil
}

func (repo *UserDatabaseRepository) FindUserByEmail(ctx context.Context, email string) (*domains.User, error) {
	var user domains.User
	err := repo.db.WithContext(ctx).Where("email = ?", email).First(&user).Error
	if err != nil {
		repo.logger.WithFields(logrus.Fields{
			"method": "FindUserByEmail",
			"email":  email,
			"error":  err,
		}).Error("Error finding user by email")

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	repo.logger.WithFields(logrus.Fields{
		"method": "FindUserByEmail",
		"email":  email,
	}).Info("User found successfully")

	return &user, nil
}

func (repo *UserDatabaseRepository) FindAllUsers(ctx context.Context) ([]domains.User, error) {
	var users []domains.User
	err := repo.db.WithContext(ctx).Find(&users).Error
	if err != nil {
		repo.logger.WithFields(logrus.Fields{
			"method": "FindAllUsers",
			"error":  err,
		}).Error("Error occurred while fetching all users")
		return nil, err
	}

	repo.logger.WithFields(logrus.Fields{
		"method": "FindAllUsers",
		"count":  len(users),
	}).Info("Users fetched successfully")

	return users, nil
}
