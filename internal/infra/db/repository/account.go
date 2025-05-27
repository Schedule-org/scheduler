package repository

import (
	"context"

	"github.com/hebertzin/scheduler/internal/domain"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type AccountDatabaseRepository struct {
	db     *gorm.DB
	logger *logrus.Logger
}

func NewAccountRepository(db *gorm.DB, logger *logrus.Logger) *AccountDatabaseRepository {
	return &AccountDatabaseRepository{
		db:     db,
		logger: logger,
	}
}

func (repo *AccountDatabaseRepository) Add(ctx context.Context, account *domain.Account) (*domain.Account, error) {
	if err := repo.db.WithContext(ctx).
		Create(account).Error; err != nil {
		return nil, err
	}
	return account, nil
}
