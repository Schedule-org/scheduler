package repository

import (
	"context"

	"github.com/hebertzin/scheduler/internal/domains"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ProfessionalsAvailabilityRepository interface {
	Add(ctx context.Context, availability *domains.ProfessionalAvailability) (*domains.ProfessionalAvailability, error)
}

type ProfessionalsAvailabilityDatabaseRepository struct {
	db     *gorm.DB
	logger *logrus.Logger
}

func NewProfessionalsAvailabilityRepository(db *gorm.DB, logger *logrus.Logger) *ProfessionalsAvailabilityDatabaseRepository {
	return &ProfessionalsAvailabilityDatabaseRepository{
		db:     db,
		logger: logger,
	}
}

func (repo *ProfessionalsAvailabilityDatabaseRepository) Add(ctx context.Context, availability *domains.ProfessionalAvailability) (*domains.ProfessionalAvailability, error) {
	if err := repo.db.WithContext(ctx).
		Create(availability).Error; err != nil {
		return nil, err
	}
	return availability, nil
}
