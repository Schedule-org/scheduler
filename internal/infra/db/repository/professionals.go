package repository

import (
	"context"

	"github.com/hebertzin/scheduler/internal/domains"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ProfessionalsRepository interface {
	Add(ctx context.Context, establishment *domains.Professionals) (*domains.Professionals, error)
	FindProfessionalById(ctx context.Context, email string) (*domains.Professionals, error)
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
