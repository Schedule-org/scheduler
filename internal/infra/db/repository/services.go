package repository

import (
	"context"

	"github.com/hebertzin/scheduler/internal/domain"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ServicesDatabaseRepository struct {
	db     *gorm.DB
	logger *logrus.Logger
}

func NewServicesRepository(db *gorm.DB, logger *logrus.Logger) *ServicesDatabaseRepository {
	return &ServicesDatabaseRepository{
		db:     db,
		logger: logger,
	}
}

func (repo *ServicesDatabaseRepository) Add(ctx context.Context, service *domain.Services) (*domain.Services, error) {
	if err := repo.db.WithContext(ctx).
		Create(service).Error; err != nil {
		return nil, err
	}
	return service, nil
}

func (repo *ServicesDatabaseRepository) FindServiceById(ctx context.Context, service_id string) (*domain.Services, error) {
	var service domain.Services
	if err := repo.db.WithContext(ctx).
		Where("id = ?", service_id).
		First(&service).Error; err != nil {
		return nil, err
	}
	return &service, nil
}

func (repo *ServicesDatabaseRepository) GetAllServicesByProfessionalId(ctx context.Context, professional_id string) ([]domain.Services, error) {
	var services []domain.Services
	if err := repo.db.WithContext(ctx).
		Where("professional_id = ?", professional_id).
		Find(&services).Error; err != nil {
		return nil, err
	}
	return services, nil
}
