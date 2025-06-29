package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

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

	srv := http.Server{
		Addr:    config.Port,
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Println("Server Shutdown:", err)
	}

	<-ctx.Done()
	log.Println("timeout of 5 seconds.")
	log.Println("Server exiting")
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
