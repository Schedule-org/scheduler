package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hebertzin/scheduler/internal/infra/config/env"
	"github.com/hebertzin/scheduler/internal/infra/config/logging"
	"github.com/hebertzin/scheduler/internal/infra/db"
	router "github.com/hebertzin/scheduler/internal/presentation/routes"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var (
	requests = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"method", "endpoint"},
	)
)

func init() {
	prometheus.MustRegister(requests)
}

func handler(c *gin.Context) {
	requests.WithLabelValues(c.Request.Method, c.FullPath()).Inc()
	c.String(http.StatusOK, "Hello, Grafana!")
}

// @title Scheduler app
// @version 1.0
// @description Saas plataforma

// @contact.name Hebert Santos
// @contact.url https://www.hebertzin.com/
// @contact.email hebertsantosdeveloper@gmail.com

// @BasePath /api/v1
func main() {
	appConfig := env.LoadConfig()
	log := logging.InitLogger()
	database := db.ConnectDatabase(appConfig)
	if err := db.Migrate(database); err != nil {
		panic("Database migration failed: " + err.Error())
	}
	r := gin.Default()
	r.GET("/api/v1/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/", handler)
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	router.UsersGroupRouter(r, database, log)
	router.EstablishmentGroupRouter(r, database, log)
	router.ProfessionalsGroupRouter(r, database, log)
	router.ServicesGroupRouter(r, database, log)
	router.ProfessionalAvailabilityGroupRouter(r, database, log)
	if err := r.Run(":8080"); err != nil {
		println("some error has been occurred:", err.Error())
	}

}
