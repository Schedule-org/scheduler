package router

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func InitRoutes(router *gin.Engine, db *gorm.DB, logger *logrus.Logger) {
	UsersGroupRouter(router, db, logger)
	EstablishmentGroupRouter(router, db, logger)
	ProfessionalsGroupRouter(router, db, logger)
	ServicesGroupRouter(router, db, logger)
	AppointmentGroupRouter(router, db, logger)
	ProfessionalAvailabilityGroupRouter(router, db, logger)
	ProfessionalAvailabilityGroupRouter(router, db, logger)
}
