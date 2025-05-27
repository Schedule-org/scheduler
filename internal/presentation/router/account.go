package router

import (
	"github.com/gin-gonic/gin"
	"github.com/hebertzin/scheduler/internal/infra/factory"
	"github.com/hebertzin/scheduler/internal/presentation/middlewares"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func UsersGroupRouter(router *gin.Engine, db *gorm.DB, logger *logrus.Logger) {
	accountFactory := factory.AccountFactory(db, logger)
	v1 := router.Group("/api/v1")
	{
		v1.GET("/accounts", accountFactory.FindAllUsers)

		v1.POST("/accounts", accountFactory.Add)

		v1.Use(middlewares.ValidateParamRequest())

		v1.GET("/accounts/:id", accountFactory.FindUserById)

		v1.GET("/accounts/:id/establishments", accountFactory.FindAllEstablishmentsByUserId)
	}
}
