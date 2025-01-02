package repository

import (
	"context"
	"errors"

	"github.com/hebertzin/scheduler/internal/domains"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type EstablishmentRepository interface {
	Add(ctx context.Context, establishment *domains.Establishment) (*domains.Establishment, error)
	FindEstablishmentById(ctx context.Context, email string) (*domains.Establishment, error)
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
	repo.logger.WithFields(logrus.Fields{
		"method":        "Add",
		"establishment": establishment,
	}).Info("Init establishment creation")

	err := repo.db.WithContext(ctx).Create(establishment).Error
	if err != nil {
		repo.logger.WithFields(logrus.Fields{
			"method": "Add",
			"error":  err,
		}).Error("Error occurred while creating establishment")
		return nil, err
	}

	repo.logger.WithFields(logrus.Fields{
		"method": "Add",
		"user":   establishment.Name,
	}).Info("establishment created successfully")

	return establishment, nil
}

func (repo *EstablishmentDatabaseRepository) FindEstablishmentById(ctx context.Context, establishment_id string) (*domains.Establishment, error) {
	var establishment domains.Establishment
	err := repo.db.WithContext(ctx).Where("id = ?", establishment_id).First(&establishment).Error
	if err != nil {
		repo.logger.WithFields(logrus.Fields{
			"method":           "FindEstablishmentById",
			"establishment_id": establishment_id,
			"error":            err,
		}).Error("Error finding  establishment by ID")

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	repo.logger.WithFields(logrus.Fields{
		"method": "FindUserById",
		"userID": establishment_id,
	}).Info("establishment found successfully")

	return &establishment, nil
}
