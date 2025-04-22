package usecases

import (
	"context"
	"testing"
	"time"

	"github.com/hebertzin/scheduler/internal/core/usecases"
	"github.com/hebertzin/scheduler/internal/domain"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockAppointmentRepository struct {
	mock.Mock
}

func (m *MockAppointmentRepository) Add(ctx context.Context, appointment *domain.Appointment) (*domain.Appointment, error) {
	args := m.Called(ctx, appointment)
	return args.Get(0).(*domain.Appointment), args.Error(1)
}

func (m *MockAppointmentRepository) GetAllAppointmentsByProfessionalId(ctx context.Context, professionalID string) ([]domain.Appointment, error) {
	return nil, nil
}

func (m *MockAppointmentRepository) GetAppointmentById(ctx context.Context, appointmentID string) (*domain.Appointment, error) {
	return nil, nil
}

func TestAppointmentUseCase_Add_Success(t *testing.T) {
	mockRepo := new(MockAppointmentRepository)
	logger := logrus.New()

	appointmentUseCase := usecases.NewAppointmentUseCase(mockRepo, logger)

	scheduledDate := time.Date(2025, 1, 15, 10, 0, 0, 0, time.UTC)
	appointment := &domain.Appointment{
		ProfessionalId: "5689c151-c968-47d7-9ff8-97a863047f19",
		ScheduledDate:  scheduledDate,
		Email:          "hebertsantosdeveloper@gmail.com",
		Phone:          "13996612070",
		ServiceId:      "5689c151-c968-47d7-9ff8-97a863047f19",
		Notes:          "Some note",
	}

	mockRepo.On("Add", mock.Anything, appointment).Return(appointment, nil)
	result, _ := appointmentUseCase.Add(context.Background(), appointment)

	assert.NotNil(t, result)
	assert.Equal(t, appointment.ProfessionalId, result.ProfessionalId)
	assert.Equal(t, appointment.ScheduledDate, result.ScheduledDate)
	assert.Equal(t, appointment.Email, result.Email)
	assert.Equal(t, appointment.Phone, result.Phone)
	assert.Equal(t, appointment.ServiceId, result.ServiceId)
	assert.Equal(t, appointment.Notes, result.Notes)

	mockRepo.AssertCalled(t, "Add", mock.Anything, appointment)
}
