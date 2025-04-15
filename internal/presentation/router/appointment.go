package router

import (
	"github.com/gin-gonic/gin"
	"github.com/hebertzin/scheduler/internal/infra/factory"
	"github.com/hebertzin/scheduler/internal/presentation/middlewares"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func AppointmentGroupRouter(router *gin.Engine, db *gorm.DB, logger *logrus.Logger) {
	appointmentFactory := factory.AppointmentFactory(db, logger)
	v1 := router.Group("/api/v1")
	{
		v1.POST("/appointments", appointmentFactory.Add)

		// here we apply the middleware to all get routes from here
		v1.Use(middlewares.ValidateParamRequest())

		v1.GET("/appointments/:id/professional", appointmentFactory.GetAllAppointmentsByProfessionalId)

		v1.GET("/appointments/:id", appointmentFactory.GetAppointmentById)
	}
}
