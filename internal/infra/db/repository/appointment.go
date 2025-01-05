package repository

import (
	"context"

	"github.com/hebertzin/scheduler/internal/domains"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type AppointmentRepository interface {
	Add(ctx context.Context, appointment *domains.Appointment) (*domains.Appointment, error)
}

type AppointmentDatabaseRepository struct {
	db     *gorm.DB
	logger *logrus.Logger
}

func NewAppointmentRepository(db *gorm.DB, logger *logrus.Logger) *AppointmentDatabaseRepository {
	return &AppointmentDatabaseRepository{
		db:     db,
		logger: logger,
	}
}

func (repo *AppointmentDatabaseRepository) Add(ctx context.Context, appointment *domains.Appointment) (*domains.Appointment, error) {
	if err := repo.db.WithContext(ctx).
		Create(appointment).Error; err != nil {
		return nil, err
	}
	return appointment, nil
}
