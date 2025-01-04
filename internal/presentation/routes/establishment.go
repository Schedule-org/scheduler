package router

import (
	"github.com/gin-gonic/gin"
	"github.com/hebertzin/scheduler/internal/infra/factory"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func EstablishmentGroupRouter(router *gin.Engine, db *gorm.DB, logger *logrus.Logger) {
	establishmentFactory := factory.EstablishmentFactory(db, logger)
	v1 := router.Group("/api/v1")
	{
		v1.POST("/establishments/", establishmentFactory.Add)
		v1.GET("/establishment_id/:establishment_id", establishmentFactory.FindEstablishmentById)
		v1.GET("/establishment/professionals/:establishment_id", establishmentFactory.FindEstablishmentById)
	}
}
