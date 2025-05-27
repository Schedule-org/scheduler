package repository

import (
	"context"

	"github.com/hebertzin/scheduler/internal/domain"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ClientDatabaseRepository struct {
	db     *gorm.DB
	logger *logrus.Logger
}

func NewClientRepository(db *gorm.DB, logger *logrus.Logger) *ClientDatabaseRepository {
	return &ClientDatabaseRepository{
		db:     db,
		logger: logger,
	}
}

func (repo *ClientDatabaseRepository) Add(ctx context.Context, Client *domain.Client) (*domain.Client, error) {
	if err := repo.db.WithContext(ctx).
		Create(Client).Error; err != nil {
		return nil, err
	}
	return Client, nil
}
