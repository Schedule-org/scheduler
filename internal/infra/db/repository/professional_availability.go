package repository

import (
	"context"

	"github.com/hebertzin/scheduler/internal/domains"
	"github.com/hebertzin/scheduler/internal/infra/db/models"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

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

func (repo *ProfessionalsAvailabilityDatabaseRepository) GetProfessionalAvailabilityById(ctx context.Context, professional_id string) ([]domains.ProfessionalAvailability, error) {
	var availability []domains.ProfessionalAvailability
	if err := repo.db.WithContext(ctx).
		Model(&models.Professional{}).
		Where("professional_id = ?", professional_id).Error; err != nil {
		return nil, err
	}
	return availability, nil
}
