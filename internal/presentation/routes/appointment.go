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
		v1.POST("/appointment/", appointmentFactory.Add)
	}
}
