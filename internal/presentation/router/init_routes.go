package router

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func StartApi(router *gin.Engine, db *gorm.DB, logger *logrus.Logger) {
	StartAccountsApi(router, db, logger)
	StartEstablishmentApi(router, db, logger)
	StartProfessionalsApi(router, db, logger)
	StartServicesApi(router, db, logger)
	StartAppointmentApi(router, db, logger)
	StartProfessionalAvailabilityApi(router, db, logger)
}
