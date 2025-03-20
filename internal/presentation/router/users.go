package router

import (
	"github.com/gin-gonic/gin"
	"github.com/hebertzin/scheduler/internal/infra/factory"
	"github.com/hebertzin/scheduler/internal/presentation/middlewares"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func UsersGroupRouter(router *gin.Engine, db *gorm.DB, logger *logrus.Logger) {
	usersFactory := factory.UsersFactory(db, logger)
	v1 := router.Group("/api/v1")
	{
		v1.GET("/users/:id", middlewares.ValidateParamRequest(), usersFactory.FindUserById)
		v1.GET("/users/:id/establishments", middlewares.ValidateParamRequest(), usersFactory.FindAllEstablishmentsByUserId)
		v1.GET("/users", usersFactory.FindAllUsers)
		v1.POST("/users", usersFactory.Add)
	}
}
