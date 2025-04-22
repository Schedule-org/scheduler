package repository

import (
	"context"

	"github.com/hebertzin/scheduler/internal/domain"
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

func (repo *ProfessionalsAvailabilityDatabaseRepository) Add(ctx context.Context, availability *domain.ProfessionalAvailability) (*domain.ProfessionalAvailability, error) {
	if err := repo.db.WithContext(ctx).
		Create(availability).Error; err != nil {
		return nil, err
	}
	return availability, nil
}

func (repo *ProfessionalsAvailabilityDatabaseRepository) GetProfessionalAvailabilityById(ctx context.Context, professional_id string) ([]domain.ProfessionalAvailability, error) {
	var availability []domain.ProfessionalAvailability
	if err := repo.db.WithContext(ctx).
		Model(&models.Professional{}).
		Where("professional_id = ?", professional_id).Error; err != nil {
		return nil, err
	}
	return availability, nil
}
