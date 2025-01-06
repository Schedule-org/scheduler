package repository

import (
	"context"

	"github.com/hebertzin/scheduler/internal/domains"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type AppointmentRepository interface {
	Add(ctx context.Context, appointment *domains.Appointment) (*domains.Appointment, error)
	GetAllAppointmentsByProfessionalId(ctx context.Context, professional_id string) ([]domains.Appointment, error)
	GetAppointmentById(ctx context.Context, appointment_id string) (*domains.Appointment, error)
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

func (repo *AppointmentDatabaseRepository) GetAllAppointmentsByProfessionalId(ctx context.Context, professional_id string) ([]domains.Appointment, error) {
	var appointments []domains.Appointment
	err := repo.db.WithContext(ctx).Find(&appointments).Error
	if err != nil {
		return nil, err
	}
	return appointments, nil
}

func (repo *AppointmentDatabaseRepository) GetAppointmentById(ctx context.Context, appointment_id string) (*domains.Appointment, error) {
	var appointment domains.Appointment
	if err := repo.db.WithContext(ctx).
		Where("id = ?", appointment_id).
		First(&appointment).Error; err != nil {
		return nil, err
	}
	return &appointment, nil
}
