package repository

import (
	"context"

	"github.com/hebertzin/scheduler/internal/domains"
	"github.com/hebertzin/scheduler/internal/infra/db/models"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type EstablishmentRepository interface {
	Add(ctx context.Context, establishment *domains.Establishment) (*domains.Establishment, error)
	GetAllProfessionalsByEstablishmentId(ctx context.Context, establishment_id string) ([]domains.Professionals, error)
	FindEstablishmentById(ctx context.Context, email string) (*domains.Establishment, error)
	GetEstablishmentReport(ctx context.Context, establishment_id string) (*domains.EstablishmentReport, error)
	UpdateEstablishmentById(ctx context.Context, establishment_id string, establishmentData *domains.Establishment) (*domains.Establishment, error)
}

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

func (repo *EstablishmentDatabaseRepository) Add(ctx context.Context, establishment *domains.Establishment) (*domains.Establishment, error) {
	if err := repo.db.WithContext(ctx).
		Create(establishment).Error; err != nil {
		return nil, err
	}
	return establishment, nil
}

func (repo *EstablishmentDatabaseRepository) FindEstablishmentById(ctx context.Context, establishment_id string) (*domains.Establishment, error) {
	var establishment domains.Establishment
	if err := repo.db.WithContext(ctx).
		Where("id = ?", establishment_id).
		First(&establishment).Error; err != nil {
		return nil, err
	}
	return &establishment, nil
}

func (repo *EstablishmentDatabaseRepository) GetAllProfessionalsByEstablishmentId(ctx context.Context, establishment_id string) ([]domains.Professionals, error) {
	var professionals []domains.Professionals
	if err := repo.db.WithContext(ctx).
		Model(&models.Professional{}).
		Where("establishment_id = ?", establishment_id).Error; err != nil {
		return nil, err
	}
	return professionals, nil
}

func (repo *EstablishmentDatabaseRepository) UpdateEstablishmentById(ctx context.Context, establishment_id string, establishmentData *domains.Establishment) (*domains.Establishment, error) {
	if err := repo.db.WithContext(ctx).
		Model(&models.Establishment{}).
		Where("id = ?", establishment_id).
		Updates(establishmentData).Error; err != nil {
		return nil, err
	}
	return establishmentData, nil
}

func (repo *EstablishmentDatabaseRepository) GetEstablishmentReport(ctx context.Context, establishment_id string) (*domains.EstablishmentReport, error) {
	var stats domains.EstablishmentReport
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
