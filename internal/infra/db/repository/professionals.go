package repository

import (
	"context"

	"github.com/hebertzin/scheduler/internal/domain"
	"github.com/hebertzin/scheduler/internal/infra/db/models"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ProfessionalsDatabaseRepository struct {
	db     *gorm.DB
	logger *logrus.Logger
}

func NewProfessionalsRepository(db *gorm.DB, logger *logrus.Logger) *ProfessionalsDatabaseRepository {
	return &ProfessionalsDatabaseRepository{
		db:     db,
		logger: logger,
	}
}

func (repo *ProfessionalsDatabaseRepository) Add(ctx context.Context, professional *domain.Professionals) (*domain.Professionals, error) {
	if err := repo.db.WithContext(ctx).
		Create(professional).Error; err != nil {
		return nil, err
	}
	return professional, nil
}

func (repo *ProfessionalsDatabaseRepository) FindProfessionalById(ctx context.Context, professional_id string) (*domain.Professionals, error) {
	var professionals domain.Professionals
	if err := repo.db.WithContext(ctx).
		Where("id = ?", professional_id).
		First(&professionals).Error; err != nil {
		return nil, err
	}
	return &professionals, nil
}

func (repo *ProfessionalsDatabaseRepository) UpdateProfessionalById(ctx context.Context, professionail_id string, professionalData *domain.Professionals) (*domain.Professionals, error) {
	if err := repo.db.WithContext(ctx).
		Model(&models.Professional{}).
		Where("id = ?", professionail_id).
		Updates(professionalData).Error; err != nil {
		return nil, err
	}
	return professionalData, nil
}
