package router

import (
	"github.com/gin-gonic/gin"
	"github.com/hebertzin/scheduler/internal/infra/factory"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func ProfessionalsGroupRouter(router *gin.Engine, db *gorm.DB, logger *logrus.Logger) {
	professionalFactory := factory.ProfessionalFactory(db, logger)
	v1 := router.Group("/api/v1")
	{
		v1.POST("/professionals/", professionalFactory.Add)
		v1.GET("/professionals/:professional_id", professionalFactory.FindEstablishmentById)
	}
}
