package router

import (
	"github.com/gin-gonic/gin"
	"github.com/hebertzin/scheduler/internal/infra/factory"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func StartServicesApi(router *gin.Engine, db *gorm.DB, logger *logrus.Logger) {
	servicesFactory := factory.ServiceFactory(db, logger)
	v1 := router.Group("/api/v1")
	{
		v1.POST("/services", servicesFactory.Add)
		v1.GET("/services/:id", servicesFactory.FindServiceById)
		v1.GET("/services/:id/all", servicesFactory.GetAllServicesByProfessionalId)
	}
}
