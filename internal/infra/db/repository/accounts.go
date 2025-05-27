package repository

import (
	"context"

	"github.com/hebertzin/scheduler/internal/domain"
	"github.com/hebertzin/scheduler/internal/infra/db/models"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type AccountDatabaseRepository struct {
	db     *gorm.DB
	logger *logrus.Logger
}

func NewAccountsRepository(db *gorm.DB, logger *logrus.Logger) *AccountDatabaseRepository {
	return &AccountDatabaseRepository{
		db:     db,
		logger: logger,
	}
}

func (repo *AccountDatabaseRepository) Add(ctx context.Context, user *domain.Account) (*domain.Account, error) {
	if err := repo.db.WithContext(ctx).Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *AccountDatabaseRepository) FindAccountById(ctx context.Context, id string) (*domain.Account, error) {
	var user *domain.Account
	if err := repo.db.WithContext(ctx).
		Model(&models.Accounts{}).
		Where("id = ?", id).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *AccountDatabaseRepository) FindAccountByEmail(ctx context.Context, email string) (*domain.Account, error) {
	var user *domain.Account
	if err := repo.db.WithContext(ctx).
		Model(&models.Accounts{}).
		Where("email = ?", email).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *AccountDatabaseRepository) FindAllAccounts(ctx context.Context) ([]domain.Account, error) {
	var users []domain.Account
	err := repo.db.WithContext(ctx).Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (repo *AccountDatabaseRepository) FindAllEstablishmentsByAccountId(ctx context.Context, user_id string) ([]domain.Establishment, error) {
	var establishments []domain.Establishment
	err := repo.db.WithContext(ctx).
		Where("user_id = ?", user_id).
		Find(&establishments).Error

	if err != nil {
		return nil, err
	}
	return establishments, nil
}
