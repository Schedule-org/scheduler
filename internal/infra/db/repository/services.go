package repository

import (
	"context"
	"errors"

	"github.com/hebertzin/scheduler/internal/domains"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ServicesRepository interface {
	Add(ctx context.Context, establishment *domains.Services) (*domains.Services, error)
	FindServiceById(ctx context.Context, service_id string) (*domains.Services, error)
}

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

func (repo *ServicesDatabaseRepository) Add(ctx context.Context, service *domains.Services) (*domains.Services, error) {
	repo.logger.WithFields(logrus.Fields{
		"method":  "Add",
		"service": service,
	}).Info("Init service creation")

	err := repo.db.WithContext(ctx).Create(service).Error
	if err != nil {
		repo.logger.WithFields(logrus.Fields{
			"method": "Add",
			"error":  err,
		}).Error("Error occurred while creating service")
		return nil, err
	}

	repo.logger.WithFields(logrus.Fields{
		"method":  "Add",
		"service": service.Name,
	}).Info("service created successfully")

	return service, nil
}

func (repo *ServicesDatabaseRepository) FindServiceById(ctx context.Context, service_id string) (*domains.Services, error) {
	var service domains.Services
	err := repo.db.WithContext(ctx).Where("id = ?", service_id).First(&service).Error
	if err != nil {
		repo.logger.WithFields(logrus.Fields{
			"method":     "FindServiceById",
			"service_id": service_id,
			"error":      err,
		}).Error("Error finding  service by ID")

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	repo.logger.WithFields(logrus.Fields{
		"method": "FindServiceById",
		"userID": service_id,
	}).Info("service found successfully")

	return &service, nil
}
