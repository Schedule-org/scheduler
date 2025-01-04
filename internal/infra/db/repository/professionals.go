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
	repo.logger.WithFields(logrus.Fields{
		"method":       "Add",
		"professional": professional,
	}).Info("Init professional creation")

	_ = repo.db.WithContext(ctx).Create(professional).Error
	return professional, nil
}

func (repo *ProfessionalsDatabaseRepository) FindProfessionalById(ctx context.Context, professional_id string) (*domains.Professionals, error) {
	var professional domains.Professionals
	err := repo.db.WithContext(ctx).Where("id = ?", professional_id).First(&professional).Error
	if err != nil {
		return nil, err
	}
	return &professional, nil
}
