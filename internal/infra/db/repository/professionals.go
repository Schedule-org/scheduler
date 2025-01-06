package repository

import (
	"context"

	"github.com/hebertzin/scheduler/internal/domains"
	"github.com/hebertzin/scheduler/internal/infra/db/models"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ProfessionalsRepository interface {
	Add(ctx context.Context, establishment *domains.Professionals) (*domains.Professionals, error)
	FindProfessionalById(ctx context.Context, email string) (*domains.Professionals, error)
	UpdateProfessionalById(ctx context.Context, professional_id string, professionalData *domains.Professionals) (*domains.Professionals, error)
}

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

func (repo *ProfessionalsDatabaseRepository) Add(ctx context.Context, professional *domains.Professionals) (*domains.Professionals, error) {
	if err := repo.db.WithContext(ctx).
		Create(professional).Error; err != nil {
		return nil, err
	}
	return professional, nil
}

func (repo *ProfessionalsDatabaseRepository) FindProfessionalById(ctx context.Context, professional_id string) (*domains.Professionals, error) {
	var professionals domains.Professionals
	if err := repo.db.WithContext(ctx).
		Where("id = ?", professional_id).
		First(&professionals).Error; err != nil {
		return nil, err
	}
	return &professionals, nil
}

func (repo *ProfessionalsDatabaseRepository) UpdateProfessionalById(ctx context.Context, professionail_id string, professionalData *domains.Professionals) (*domains.Professionals, error) {
	if err := repo.db.WithContext(ctx).
		Model(&models.Professional{}).
		Where("id = ?", professionail_id).
		Updates(professionalData).Error; err != nil {
		return nil, err
	}
	return professionalData, nil
}
