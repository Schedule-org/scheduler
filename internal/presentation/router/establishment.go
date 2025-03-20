package router

import (
	"github.com/gin-gonic/gin"
	"github.com/hebertzin/scheduler/internal/infra/factory"
	"github.com/hebertzin/scheduler/internal/presentation/middlewares"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func EstablishmentGroupRouter(router *gin.Engine, db *gorm.DB, logger *logrus.Logger) {
	establishmentFactory := factory.EstablishmentFactory(db, logger)
	v1 := router.Group("/api/v1")
	{
		v1.GET("/establishments/:id", middlewares.ValidateParamRequest(), establishmentFactory.FindEstablishmentById)
		v1.GET("/establishments/:id/professionals", middlewares.ValidateParamRequest(), establishmentFactory.GetAllProfessinalsByEstablishmentId)
		v1.GET("/establishments/:id/report", middlewares.ValidateParamRequest(), establishmentFactory.GetEstablishmentReport)
		v1.POST("/establishments", establishmentFactory.Add)
		v1.PUT("/establishments/:id/update", middlewares.ValidateParamRequest(), establishmentFactory.UpdateEstablishmentById)
	}
}
