package repository

import (
	"context"

	"github.com/hebertzin/scheduler/internal/domain"
	"github.com/hebertzin/scheduler/internal/infra/db/models"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type EstablishmentDatabaseRepository struct {
	db     *gorm.DB
	logger *logrus.Logger
}

func NewEstablishmentRepository(db *gorm.DB, logger *logrus.Logger) *EstablishmentDatabaseRepository {
	return &EstablishmentDatabaseRepository{
		db:     db,
		logger: logger,
	}
}

func (repo *EstablishmentDatabaseRepository) Add(ctx context.Context, establishment *domain.Establishment) (*domain.Establishment, error) {
	if err := repo.db.WithContext(ctx).
		Create(establishment).Error; err != nil {
		return nil, err
	}
	return establishment, nil
}

func (repo *EstablishmentDatabaseRepository) FindEstablishmentById(ctx context.Context, establishment_id string) (*domain.Establishment, error) {
	var establishment domain.Establishment
	if err := repo.db.WithContext(ctx).
		Where("id = ?", establishment_id).
		First(&establishment).Error; err != nil {
		return nil, err
	}
	return &establishment, nil
}

func (repo *EstablishmentDatabaseRepository) GetAllProfessionalsByEstablishmentId(ctx context.Context, establishment_id string) ([]domain.Professionals, error) {
	var professionals []domain.Professionals
	if err := repo.db.WithContext(ctx).
		Model(&models.Professional{}).
		Where("establishment_id = ?", establishment_id).Error; err != nil {
		return nil, err
	}
	return professionals, nil
}

func (repo *EstablishmentDatabaseRepository) UpdateEstablishmentById(ctx context.Context, establishment_id string, establishmentData *domain.Establishment) (*domain.Establishment, error) {
	if err := repo.db.WithContext(ctx).
		Model(&models.Establishment{}).
		Where("id = ?", establishment_id).
		Updates(establishmentData).Error; err != nil {
		return nil, err
	}
	return establishmentData, nil
}

func (repo *EstablishmentDatabaseRepository) GetEstablishmentReport(ctx context.Context, establishment_id string) (*domain.EstablishmentReport, error) {
	var stats domain.EstablishmentReport
	if err := repo.db.WithContext(ctx).
		Model(&models.Professional{}).
		Where("establishment_id = ?", establishment_id).
		Count(&stats.TotalProfessionals).Error; err != nil {
		return nil, err
	}

	if err := repo.db.WithContext(ctx).
		Model(&models.Services{}).
		Where("establishment_id = ?", establishment_id).
		Count(&stats.TotalServices).Error; err != nil {
		return nil, err
	}

	return &stats, nil
}
