package router

import (
	"github.com/gin-gonic/gin"
	"github.com/hebertzin/scheduler/internal/infra/factory"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func AppointmentGroupRouter(router *gin.Engine, db *gorm.DB, logger *logrus.Logger) {
	appointmentFactory := factory.AppointmentFactory(db, logger)
	v1 := router.Group("/api/v1")
	{
		v1.GET("/appointment/:professional_id/all", appointmentFactory.GetAllAppointmentsByProfessionalId)
		v1.GET("/appointment/:appointment_id", appointmentFactory.GetAppointmentById)
		v1.POST("/appointment", appointmentFactory.Add)
	}
}
