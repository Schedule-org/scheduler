package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hebertzin/scheduler/internal/domain"
	"github.com/hebertzin/scheduler/internal/infra/config/env"
	"github.com/hebertzin/scheduler/internal/infra/config/logging"
	"github.com/hebertzin/scheduler/internal/infra/db"
	"github.com/hebertzin/scheduler/internal/presentation/router"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Scheduler app
// @version 1.0
// @description Saas plataform

// @contact.name Hebert Santos
// @contact.url https://www.hebertzin.com/
// @contact.email hebertsantosdeveloper@gmail.com

// @BasePath /api/v1
func main() {
	config, _ := env.LoadConfiguration("/config/dev.json")
	database := db.ConnectDatabase(config)
	r := createRouter()
	configureSwagger(config, r)
	configureMigration(config, database)
	cofigureMetrics(config, r)
	loggging := configureLogging(config)
	router.StartApi(r, database, loggging)
	if err := r.Run(config.Port); err != nil {
		println("some error has been occurred:", err.Error())
	}
}

func createRouter() *gin.Engine {
	return gin.Default()
}

func configureMigration(config *domain.ServiceConfig, database *gorm.DB) {
	if config.RunMigrationEnabled {
		if err := db.Migrate(database); err != nil {
			panic("Database migration failed: " + err.Error())
		}
	}
}

func configureSwagger(config *domain.ServiceConfig, router *gin.Engine) {
	if config.SwaggerEnabled {
		router.GET("/api/v1/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
}

func grafanaHandler(c *gin.Context) {
	var requests *prometheus.CounterVec
	requests.WithLabelValues(c.Request.Method, c.FullPath()).Inc()
}

func cofigureMetrics(config *domain.ServiceConfig, router *gin.Engine) {
	if config.GrafanaEnabled {
		requests := prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name: "http_requests_total",
				Help: "Total number of HTTP requests",
			},
			[]string{"method", "endpoint"},
		)
		prometheus.MustRegister(requests)

		router.GET("/grafana", grafanaHandler)
		router.GET("/metrics", gin.WrapH(promhttp.Handler()))
	}
}

func configureLogging(config *domain.ServiceConfig) *logrus.Logger {
	if config.LoggingEnabled {
		return logging.InitLogger()
	}
	return nil
}
